package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPatron_ReturnBook(t *testing.T) {
	t.Run("return a normal book", func(t *testing.T) {
		t.Parallel()
		patron := newExamplePatron(t, domain.PatronTypeRegular, nil, nil)
		bookID := domain.BookID(uuid.NewString())

		patron.ReturnBook(bookID)
	})
	t.Run("return a overdue checkout book", func(t *testing.T) {
		bookInfo := newExampleBookInformation(t, domain.BookTypeCirculating)
		patron := newExamplePatron(t, domain.PatronTypeRegular, nil, map[domain.LibraryBranchID][]domain.BookID{
			bookInfo.PlacedAt: {bookInfo.BookID},
		})

		patron.ReturnBook(bookInfo.BookID)

		assert.Len(t, patron.OverdueCheckouts()[bookInfo.PlacedAt], 0)
	})
}
