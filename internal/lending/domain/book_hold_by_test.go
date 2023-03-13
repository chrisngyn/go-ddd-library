package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestBook_HoldBy(t *testing.T) {
	patronID := domain.PatronID("patron1")
	t.Run("error when book is not available", func(t *testing.T) {
		t.Parallel()
		book := newExampleBookOnHold(t)
		holdDuration := newExampleCloseEndedHoldDuration(t)

		err := book.HoldBy(patronID, holdDuration)

		assert.ErrorIs(t, err, domain.ErrBookNotAvailable)
	})
	t.Run("hold success", func(t *testing.T) {
		t.Parallel()
		book := newExampleAvailableBook(t)
		holdDuration := newExampleCloseEndedHoldDuration(t)

		err := book.HoldBy(patronID, holdDuration)
		require.NoError(t, err)

		assert.Equal(t, domain.BookStatusOnHold, book.Status())
		assert.Equal(t, patronID, book.BookHoldInfo().ByPatron)
		assert.Equal(t, holdDuration.Till(), book.BookHoldInfo().Till)
	})
}
