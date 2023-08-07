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
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPostgresBookRepository_Update(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := adapters.NewPostgresBookRepository(db)
	dbBook := addExampleAvailableBook(t, db)

	var updatedBook *book.Book

	err := repo.Update(context.Background(), uuid.MustParse(dbBook.ID), func(ctx context.Context, book *book.Book) error {
		holdDuration, err := patron.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = book.HoldBy(uuid.New(), holdDuration.Till())
		require.NoError(t, err)

		updatedBook = book
		return nil
	})
	require.NoError(t, err)

	assertPersistedBookEquals(t, repo, updatedBook)
}

func addExampleAvailableBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	aBook := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusAvailable,
	}
	err := aBook.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a available book: %v", err)
	}
	if err := aBook.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload book: %v", err)
	}
	return aBook
}

func assertPersistedBookEquals(t *testing.T, repo adapters.PostgresBookRepository, aBook *book.Book) {
	t.Helper()
	persistedBook, err := repo.Get(context.Background(), aBook.ID())
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			book.Type{},
			book.Status{},
			book.Book{},
		),
	}

	assert.True(
		t,
		cmp.Equal(aBook, &persistedBook, cmpOpts...),
		cmp.Diff(aBook, &persistedBook, cmpOpts...),
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
	err := patronRepo.UpdateWithBook(ctx, uuid.MustParse(dbPatron.ID), uuid.MustParse(dbBook.ID), func(ctx context.Context, aPatron *patron.Patron, aBook *book.Book) error {
		holdDuration, err := patron.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = aPatron.PlaceOnHold(aBook.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = aBook.HoldBy(aPatron.ID(), holdDuration.Till())
		require.NoError(t, err)

		return nil
	})
	require.NoError(t, err)

	// Patron checkout book
	var updatedPatron *patron.Patron
	var updatedBook *book.Book
	err = repo.UpdateWithPatron(ctx, uuid.MustParse(dbBook.ID), func(ctx context.Context, aBook *book.Book, aPatron *patron.Patron) error {
		err := aBook.Checkout(aPatron.ID(), time.Now())
		require.NoError(t, err)

		err = aPatron.Checkout(aBook.ID())
		require.NoError(t, err)

		updatedPatron = aPatron
		updatedBook = aBook

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
			BookID:          uuid.MustParse(dbBook.ID),
			LibraryBranchID: uuid.MustParse(dbBook.LibraryBranchID),
			PatronID:        uuid.MustParse(dbBook.PatronID.String),
			HoldTill:        dbBook.HoldTill.Time,
		},
	}

	actual, err := repo.ListExpiredHolds(ctx, time.Now())
	require.NoError(t, err)

	assert.ElementsMatch(t, expectedQueryExpiredHolds, actual)

}

func addExpiredHoldBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	aBook := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusOnHold,
		PatronID:        null.StringFrom(uuid.NewString()),
		HoldTill:        null.TimeFrom(time.Now().AddDate(0, 0, -1)),
	}
	err := aBook.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a expried book: %v", err)
	}
	if err := aBook.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload expried book: %v", err)
	}
	return aBook
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
			PatronID:        uuid.MustParse(dbBook.PatronID.String),
			BookID:          uuid.MustParse(dbBook.ID),
			LibraryBranchID: uuid.MustParse(dbBook.LibraryBranchID),
		},
	}

	actual, err := repo.ListOverdueCheckouts(ctx, time.Now(), 10)
	require.NoError(t, err)

	assert.ElementsMatch(t, expectOverdueCheckouts, actual)
}

func addOverdueCheckoutBook(t *testing.T, db *sql.DB) *models.Book {
	t.Helper()
	aBook := &models.Book{
		ID:              uuid.NewString(),
		LibraryBranchID: uuid.NewString(),
		BookType:        models.BookTypeCirculating,
		BookStatus:      models.BookStatusCheckedOut,
		PatronID:        null.StringFrom(uuid.NewString()),
		CheckedOutAt:    null.TimeFrom(time.Now().AddDate(0, 0, -70)), // just a day so far
	}
	err := aBook.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		t.Fatalf("Cannot create a expried book: %v", err)
	}
	if err := aBook.Reload(context.Background(), db); err != nil {
		t.Fatalf("Cannot reload expried book: %v", err)
	}
	return aBook
}
