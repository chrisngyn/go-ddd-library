package patron_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPatron_MarkOverdueCheckout(t *testing.T) {
	t.Run("mark overdue checkout successfully", func(t *testing.T) {
		t.Parallel()
		aPatron := newExamplePatron(t, patron.TypeRegular, nil, nil)
		bookInfo := newExampleBookInformation(t, book.TypeCirculating)

		aPatron.MarkOverdueCheckout(bookInfo.BookID, bookInfo.PlacedAt)

		overdueCheckouts := aPatron.OverdueCheckouts()
		assert.Equal(t, 1, aPatron.OverdueCheckouts().TotalAt(bookInfo.PlacedAt))
		assert.Equal(t, bookInfo.BookID, overdueCheckouts[bookInfo.PlacedAt][0])
	})
}
