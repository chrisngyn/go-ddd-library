//go:build integration

package adapters_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

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

	assert.EqualValues(t, book.ID(), persistedBook.ID())
	assert.EqualValues(t, book.Status(), persistedBook.Status())
	assert.EqualValues(t, book.BookInfo(), persistedBook.BookInfo())
	assert.EqualValues(t, book.BookHoldInfo().ByPatron, persistedBook.BookHoldInfo().ByPatron)
	assert.EqualValues(t, book.BookHoldInfo().Till.Unix(), persistedBook.BookHoldInfo().Till.Unix())
	assert.EqualValues(t, book.BookCheckedOutInfo().ByPatron, persistedBook.BookCheckedOutInfo().ByPatron)
	assert.EqualValues(t, book.BookCheckedOutInfo().At.Unix(), persistedBook.BookCheckedOutInfo().At.Unix())
}
