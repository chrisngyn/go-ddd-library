package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestBook_CancelHold(t *testing.T) {
	t.Run("error when cancel a book not on hold", func(t *testing.T) {
		t.Parallel()
		book := newExampleCheckedOutBook(t)

		err := book.CancelHold()

		assert.ErrorIs(t, err, domain.ErrBookNotOnHold)
	})

	t.Run("cancel hold a book on hold", func(t *testing.T) {
		t.Parallel()
		book := newExampleBookOnHold(t)

		err := book.CancelHold()
		require.NoError(t, err)

		assert.Equal(t, domain.BookStatusAvailable, book.Status())
		assert.True(t, book.BookHoldInfo().IsZero())
	})
}
