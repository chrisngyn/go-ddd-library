package book_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestBook_CancelHold(t *testing.T) {
	t.Run("error when cancel a book not on hold", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleCheckedOutBook(t)

		err := aBook.CancelHold()

		assert.ErrorIs(t, err, book.ErrBookNotOnHold)
	})

	t.Run("cancel hold a book on hold", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleBookOnHold(t)

		err := aBook.CancelHold()
		require.NoError(t, err)

		assert.Equal(t, book.StatusAvailable, aBook.Status())
		assert.True(t, aBook.BookHoldInfo().IsZero())
	})
}
