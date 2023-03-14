//go:build integration

package adapters_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPostgresBookRepository_Update(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := adapters.NewPostgresBookRepository(db)
	dbBook := addExampleAvailableBook(t, db)

	var updatedBook *domain.Book

	err := repo.Update(context.Background(), domain.BookID(dbBook.ID), func(ctx context.Context, book *domain.Book) error {
		holdDuration, err := domain.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = book.HoldBy(domain.PatronID(uuid.NewString()), holdDuration)
		require.NoError(t, err)

		updatedBook = book
		return nil
	})
	require.NoError(t, err)

	assertPersistedBookEquals(t, repo, updatedBook)
}

func addExampleAvailableBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	book := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusAvailable,
	}
	err := book.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a available book: %v", err)
	}
	if err := book.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload book: %v", err)
	}
	return book
}

func assertPersistedBookEquals(t *testing.T, repo adapters.PostgresBookRepository, book *domain.Book) {
	t.Helper()
	persistedBook, err := repo.Get(context.Background(), book.ID())
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			domain.BookType{},
			domain.BookStatus{},
			domain.Book{},
		),
	}

	assert.True(
		t,
		cmp.Equal(book, &persistedBook, cmpOpts...),
		cmp.Diff(book, &persistedBook, cmpOpts...),
	)

}

func TestPostgresBookRepository_UpdateWithPatron(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})
	ctx := context.Background()

	repo := adapters.NewPostgresBookRepository(db)
	patronRepo := adapters.NewPostgresPatronRepository(db)

	dbPatron := addExamplePatron(t, db)
	dbBook := addExampleAvailableBook(t, db)

	// Patron hold a book
	err := patronRepo.UpdateWithBook(ctx, domain.PatronID(dbPatron.ID), domain.BookID(dbBook.ID), func(ctx context.Context, patron *domain.Patron, book *domain.Book) error {
		holdDuration, err := domain.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = patron.PlaceOnHold(book.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = book.HoldBy(patron.ID(), holdDuration)
		require.NoError(t, err)

		return nil
	})
	require.NoError(t, err)

	// Patron checkout book
	var updatedPatron *domain.Patron
	var updatedBook *domain.Book
	err = repo.UpdateWithPatron(ctx, domain.BookID(dbBook.ID), func(ctx context.Context, book *domain.Book, patron *domain.Patron) error {
		err := book.Checkout(patron.ID(), time.Now())
		require.NoError(t, err)

		err = patron.Checkout(book.ID())
		require.NoError(t, err)

		updatedPatron = patron
		updatedBook = book

		return nil
	})
	require.NoError(t, err)

	assertPersistedBookEquals(t, repo, updatedBook)
	assertPersistedPatronEquals(t, patronRepo, updatedPatron)
}
