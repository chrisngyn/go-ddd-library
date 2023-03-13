package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPatron_MarkOverdueCheckout(t *testing.T) {
	t.Run("mark overdue checkout successfully", func(t *testing.T) {
		t.Parallel()
		patron := newExamplePatron(t, domain.PatronTypeRegular, nil, nil)
		bookInfo := newExampleBookInformation(t, domain.BookTypeCirculating)

		patron.MarkOverdueCheckout(bookInfo)

		overdueCheckouts := patron.OverdueCheckouts()
		assert.Len(t, overdueCheckouts[bookInfo.PlacedAt], 1)
		assert.Equal(t, bookInfo.BookID, overdueCheckouts[bookInfo.PlacedAt][0])
	})
}
