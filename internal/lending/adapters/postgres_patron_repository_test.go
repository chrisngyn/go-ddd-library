//go:build integration

package adapters_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
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

	assert.EqualValues(t, patron.ID(), persistedPatron.ID())
	assert.EqualValues(t, patron.PatronType(), persistedPatron.PatronType())
	assert.EqualValues(t, patron.Holds(), persistedPatron.Holds())
	assert.EqualValues(t, patron.OverdueCheckouts(), persistedPatron.OverdueCheckouts())
}
