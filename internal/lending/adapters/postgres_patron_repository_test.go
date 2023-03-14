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
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPostgresPatronRepository_Update(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := adapters.NewPostgresPatronRepository(db)
	dbPatron := addExamplePatron(t, db)

	var updatedPatron *domain.Patron

	err := repo.Update(context.Background(), domain.PatronID(dbPatron.ID), func(ctx context.Context, patron *domain.Patron) error {
		patron.MarkOverdueCheckout(domain.BookID(uuid.NewString()), domain.LibraryBranchID(uuid.NewString()))
		updatedPatron = patron
		return nil
	})
	require.NoError(t, err)

	assertPersistedPatronEquals(t, repo, updatedPatron)
}

func addExamplePatron(t *testing.T, db *sql.DB) *models.Patron {
	t.Helper()
	patron := &models.Patron{
		ID:         uuid.NewString(),
		PatronType: models.PatronTypeRegular,
	}
	if err := patron.Insert(context.Background(), db, boil.Infer()); err != nil {
		t.Fatalf("Error crate new patron: %v", err)
	}
	if err := patron.Reload(context.Background(), db); err != nil {
		t.Fatalf("Error to reload patron: %v", err)
	}
	return patron
}

func assertPersistedPatronEquals(t *testing.T, repo adapters.PostgresPatronRepository, patron *domain.Patron) {
	t.Helper()
	persistedPatron, err := repo.Get(context.Background(), patron.ID())
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			domain.PatronType{},
			domain.Patron{},
			domain.Hold{},
			domain.HoldDuration{},
		),
	}

	assert.True(
		t,
		cmp.Equal(patron, &persistedPatron, cmpOpts...),
		cmp.Diff(patron, &persistedPatron, cmpOpts...),
	)
}

func TestPostgresPatronRepository_UpdateWithBook(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := adapters.NewPostgresPatronRepository(db)
	dbPatron := addExamplePatron(t, db)
	dbBook := addExampleAvailableBook(t, db)

	var updatedPatron *domain.Patron
	var updatedBook *domain.Book

	err := repo.UpdateWithBook(context.Background(), domain.PatronID(dbPatron.ID), domain.BookID(dbBook.ID), func(ctx context.Context, patron *domain.Patron, book *domain.Book) error {
		holdDuration, err := domain.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = patron.PlaceOnHold(book.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = book.HoldBy(patron.ID(), holdDuration)
		require.NoError(t, err)

		updatedPatron = patron
		updatedBook = book
		return nil
	})
	require.NoError(t, err)

	assertPersistedPatronEquals(t, repo, updatedPatron)
	assertPersistedBookEquals(t, adapters.NewPostgresBookRepository(db), updatedBook)
}

func TestPostgresPatronRepository_GetPatronProfile(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})
	ctx := context.Background()

	repo := adapters.NewPostgresPatronRepository(db)
	dbPatron := addExamplePatron(t, db)
	dbBook1 := addExampleAvailableBook(t, db)
	holdDuration, err := domain.NewHoldDuration(time.Now(), 5)
	require.NoError(t, err)

	err = repo.UpdateWithBook(ctx, domain.PatronID(dbPatron.ID), domain.BookID(dbBook1.ID), func(ctx context.Context, patron *domain.Patron, book *domain.Book) error {
		err = patron.PlaceOnHold(book.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = book.HoldBy(patron.ID(), holdDuration)
		require.NoError(t, err)

		return nil
	})

	expectedQueryPatronProfile := query.PatronProfile{
		PatronID:   dbPatron.ID,
		PatronType: domain.PatronTypeRegular,
		Holds: []domain.Hold{
			{
				BookID:       domain.BookID(dbBook1.ID),
				PlacedAt:     domain.LibraryBranchID(dbBook1.LibraryBranchID),
				HoldDuration: holdDuration,
			},
		},
		CheckedOuts:      []query.CheckedOut{},
		OverdueCheckouts: []query.OverdueCheckout{},
	}
	require.NoError(t, err)

	actual, err := repo.GetPatronProfile(ctx, domain.PatronID(dbPatron.ID))
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			domain.PatronType{},
			domain.Patron{},
			domain.Hold{},
			domain.HoldDuration{},
			domain.Book{},
		),
	}

	assert.True(t,
		cmp.Equal(expectedQueryPatronProfile, actual, cmpOpts...),
		cmp.Diff(expectedQueryPatronProfile, actual, cmpOpts...),
	)
}
