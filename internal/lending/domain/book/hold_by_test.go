package book_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestBook_HoldBy(t *testing.T) {
	patronID := uuid.New()
	t.Run("error when book is not available", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleBookOnHold(t)
		err := aBook.HoldBy(patronID, time.Now().AddDate(0, 0, 10))

		assert.ErrorIs(t, err, book.ErrBookNotAvailable)
	})
	t.Run("hold success", func(t *testing.T) {
		t.Parallel()
		aBook := newExampleAvailableBook(t)
		till := time.Now().AddDate(0, 0, 10)

		err := aBook.HoldBy(patronID, till)
		require.NoError(t, err)

		assert.Equal(t, book.StatusOnHold, aBook.Status())
		assert.Equal(t, patronID, aBook.BookHoldInfo().ByPatron)
		assert.Equal(t, till, aBook.BookHoldInfo().Till)
	})
}
