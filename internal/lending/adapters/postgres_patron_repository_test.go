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
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPostgresPatronRepository_Update(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	repo := adapters.NewPostgresPatronRepository(db)
	dbPatron := addExamplePatron(t, db)

	var updatedPatron *patron.Patron

	err := repo.Update(context.Background(), uuid.MustParse(dbPatron.ID), func(ctx context.Context, patron *patron.Patron) error {
		patron.MarkOverdueCheckout(uuid.New(), uuid.New())
		updatedPatron = patron
		return nil
	})
	require.NoError(t, err)

	assertPersistedPatronEquals(t, repo, updatedPatron)
}

func addExamplePatron(t *testing.T, db *sql.DB) *models.Patron {
	t.Helper()
	aPatron := &models.Patron{
		ID:         uuid.NewString(),
		PatronType: models.PatronTypeRegular,
	}
	if err := aPatron.Insert(context.Background(), db, boil.Infer()); err != nil {
		t.Fatalf("Error crate new patron: %v", err)
	}
	if err := aPatron.Reload(context.Background(), db); err != nil {
		t.Fatalf("Error to reload patron: %v", err)
	}
	return aPatron
}

func assertPersistedPatronEquals(t *testing.T, repo adapters.PostgresPatronRepository, aPatron *patron.Patron) {
	t.Helper()
	persistedPatron, err := repo.Get(context.Background(), aPatron.ID())
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			patron.Type{},
			patron.Patron{},
			patron.Hold{},
			patron.HoldDuration{},
		),
	}

	assert.True(
		t,
		cmp.Equal(aPatron, &persistedPatron, cmpOpts...),
		cmp.Diff(aPatron, &persistedPatron, cmpOpts...),
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

	var updatedPatron *patron.Patron
	var updatedBook *book.Book

	err := repo.UpdateWithBook(context.Background(), uuid.MustParse(dbPatron.ID), uuid.MustParse(dbBook.ID), func(ctx context.Context, aPatron *patron.Patron, aBook *book.Book) error {
		holdDuration, err := patron.NewHoldDuration(time.Now(), 5)
		require.NoError(t, err)

		err = aPatron.PlaceOnHold(aBook.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = aBook.HoldBy(aPatron.ID(), holdDuration.Till())
		require.NoError(t, err)

		updatedPatron = aPatron
		updatedBook = aBook
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
	holdDuration, err := patron.NewHoldDuration(time.Now(), 5)
	require.NoError(t, err)

	err = repo.UpdateWithBook(ctx, uuid.MustParse(dbPatron.ID), uuid.MustParse(dbBook1.ID), func(ctx context.Context, aPatron *patron.Patron, aBook *book.Book) error {
		err = aPatron.PlaceOnHold(aBook.BookInfo(), holdDuration)
		require.NoError(t, err)

		err = aBook.HoldBy(aPatron.ID(), holdDuration.Till())
		require.NoError(t, err)

		return nil
	})

	expectedQueryPatronProfile := query.PatronProfile{
		PatronID:   uuid.MustParse(dbPatron.ID),
		PatronType: patron.TypeRegular,
		Holds: []patron.Hold{
			{
				BookID:       uuid.MustParse(dbBook1.ID),
				PlacedAt:     uuid.MustParse(dbBook1.LibraryBranchID),
				HoldDuration: holdDuration,
			},
		},
		CheckedOuts:      []query.CheckedOut{},
		OverdueCheckouts: []query.OverdueCheckout{},
	}
	require.NoError(t, err)

	actual, err := repo.GetPatronProfile(ctx, uuid.MustParse(dbPatron.ID))
	require.NoError(t, err)

	cmpOpts := []cmp.Option{
		cmp.AllowUnexported(
			time.Time{},
			patron.Type{},
			patron.Patron{},
			patron.Hold{},
			patron.HoldDuration{},
			book.Book{},
		),
	}

	assert.True(t,
		cmp.Equal(expectedQueryPatronProfile, actual, cmpOpts...),
		cmp.Diff(expectedQueryPatronProfile, actual, cmpOpts...),
	)
}
