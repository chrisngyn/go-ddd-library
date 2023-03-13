package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestBook_CheckIn(t *testing.T) {
	t.Run("error when check in book that is not checked out", func(t *testing.T) {
		t.Parallel()
		book := newExampleAvailableBook(t)

		err := book.CheckIn()

		assert.ErrorIs(t, err, domain.ErrBookNotCheckedOut)
	})

	t.Run("check in a checked out book", func(t *testing.T) {
		t.Parallel()
		book := newExampleCheckedOutBook(t)

		err := book.CheckIn()
		require.NoError(t, err)

		assert.Equal(t, domain.BookStatusAvailable, book.Status())
		assert.True(t, book.BookCheckedOutInfo().IsZero())
	})
}
