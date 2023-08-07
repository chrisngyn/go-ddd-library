package book_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestBook_Checkout(t *testing.T) {
	t.Run("error when book not on hold", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleAvailableBook(t)

		err := aBook.Checkout(uuid.New(), time.Now())

		assert.ErrorIs(t, err, book.ErrBookNotOnHold)
	})

	t.Run("error when checkout by different patron", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleBookOnHold(t)

		err := aBook.Checkout(uuid.New(), time.Now())

		assert.ErrorIs(t, err, book.ErrBookNotHoldByPatron)
	})

	t.Run("checkout success", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleBookOnHold(t)
		patronID := aBook.ByPatronID()
		checkedOutAt := time.Now()

		err := aBook.Checkout(patronID, checkedOutAt)
		require.NoError(t, err)

		assert.Equal(t, book.StatusCheckedOut, aBook.Status())
		assert.Equal(t, patronID, aBook.BookCheckedOutInfo().ByPatron)
		assert.Equal(t, checkedOutAt, aBook.BookCheckedOutInfo().At)
	})
}
