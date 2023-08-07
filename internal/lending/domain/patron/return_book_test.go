package patron_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPatron_ReturnBook(t *testing.T) {
	t.Run("return a normal book", func(t *testing.T) {
		t.Parallel()
		aPatron := newExamplePatron(t, patron.TypeRegular, nil, nil)
		bookID := uuid.New()

		aPatron.ReturnBook(bookID)
	})
	t.Run("return a overdue checkout book", func(t *testing.T) {
		bookInfo := newExampleBookInformation(t, book.TypeCirculating)
		aPatron := newExamplePatron(t, patron.TypeRegular, nil, map[uuid.UUID][]uuid.UUID{
			bookInfo.PlacedAt: {bookInfo.BookID},
		})

		aPatron.ReturnBook(bookInfo.BookID)

		assert.Zero(t, aPatron.OverdueCheckouts().TotalAt(bookInfo.PlacedAt))
	})
}
