package book_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestBook_CheckIn(t *testing.T) {
	t.Run("error when check in book that is not checked out", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleAvailableBook(t)

		err := aBook.CheckIn()

		assert.ErrorIs(t, err, book.ErrBookNotCheckedOut)
	})

	t.Run("check in a checked out book", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleCheckedOutBook(t)

		err := aBook.CheckIn()
		require.NoError(t, err)

		assert.Equal(t, book.StatusAvailable, aBook.Status())
		assert.True(t, aBook.BookCheckedOutInfo().IsZero())
	})
}
