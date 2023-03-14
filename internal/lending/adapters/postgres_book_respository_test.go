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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/common/tests"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
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

func TestPostgresBookRepository_ListExpiredHolds(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})
	ctx := context.Background()
	tests.TruncateTables(t, db, models.TableNames.Books)

	repo := adapters.NewPostgresBookRepository(db)
	dbBook := addExpiredHoldBook(t, db)

	expectedQueryExpiredHolds := []query.ExpiredHold{
		{
			BookID:          domain.BookID(dbBook.ID),
			LibraryBranchID: domain.LibraryBranchID(dbBook.LibraryBranchID),
			PatronID:        domain.PatronID(dbBook.PatronID.String),
			HoldTill:        dbBook.HoldTill.Time,
		},
	}

	actual, err := repo.ListExpiredHolds(ctx, time.Now())
	require.NoError(t, err)

	assert.ElementsMatch(t, expectedQueryExpiredHolds, actual)

}

func addExpiredHoldBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	book := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusOnHold,
		PatronID:        null.StringFrom(uuid.NewString()),
		HoldTill:        null.TimeFrom(time.Now().AddDate(0, 0, -1)),
	}
	err := book.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a expried book: %v", err)
	}
	if err := book.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload expried book: %v", err)
	}
	return book
}

func TestPostgresBookRepository_ListOverdueCheckouts(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})
	ctx := context.Background()
	tests.TruncateTables(t, db, models.TableNames.Books)

	repo := adapters.NewPostgresBookRepository(db)
	dbBook := addOverdueCheckoutBook(t, db)
	expectOverdueCheckouts := []query.OverdueCheckout{
		{
			PatronID:        domain.PatronID(dbBook.PatronID.String),
			BookID:          dbBook.ID,
			LibraryBranchID: dbBook.LibraryBranchID,
		},
	}

	actual, err := repo.ListOverdueCheckouts(ctx, time.Now(), 10)
	require.NoError(t, err)

	assert.ElementsMatch(t, expectOverdueCheckouts, actual)
}

func addOverdueCheckoutBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	book := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusCheckedOut,
		PatronID:        null.StringFrom(uuid.NewString()),
		CheckedOutAt:    null.TimeFrom(time.Now().AddDate(0, 0, -70)), // just a day so far
	}
	err := book.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a expried book: %v", err)
	}
	if err := book.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload expried book: %v", err)
	}
	return book
}
