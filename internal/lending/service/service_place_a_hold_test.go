//go:build component

package service_test

import (
	"context"
	"database/sql"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/common/tests"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
)

func TestServer_PlaceOnHold(t *testing.T) {
	db := database.NewSqlDB()
	t.Cleanup(func() {
		_ = db.Close()
	})

	client := tests.NewLendingHTTPClient(t, httpAddress, parentRoute)
	dbPatron := addExamplePatron(t, db)
	dbBook := addExampleAvailableBook(t, db)

	statusCode := client.PlaceOnHold(t, dbPatron.ID, dbBook.ID, 5)

	assert.EqualValues(t, http.StatusOK, statusCode)
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
