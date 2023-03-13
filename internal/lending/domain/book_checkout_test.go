package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestBook_Checkout(t *testing.T) {
	t.Run("error when book not on hold", func(t *testing.T) {
		t.Parallel()
		book := newExampleAvailableBook(t)

		err := book.Checkout(domain.PatronID(uuid.NewString()), time.Now())

		assert.ErrorIs(t, err, domain.ErrBookNotOnHold)
	})

	t.Run("error when checkout by different patron", func(t *testing.T) {
		t.Parallel()
		book := newExampleBookOnHold(t)

		err := book.Checkout(domain.PatronID(uuid.NewString()), time.Now())

		assert.ErrorIs(t, err, domain.ErrBookNotHoldByPatron)
	})

	t.Run("checkout success", func(t *testing.T) {
		t.Parallel()
		book := newExampleBookOnHold(t)
		patronID := book.ByPatronID()
		checkedOutAt := time.Now()

		err := book.Checkout(patronID, checkedOutAt)
		require.NoError(t, err)

		assert.Equal(t, domain.BookStatusCheckedOut, book.Status())
		assert.Equal(t, patronID, book.BookCheckedOutInfo().ByPatron)
		assert.Equal(t, checkedOutAt, book.BookCheckedOutInfo().At)
	})
}
