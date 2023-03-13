package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestPatron_CancelHold(t *testing.T) {
	t.Run("cancel a hold that not existed", func(t *testing.T) {
		t.Parallel()
		patron := newExamplePatron(t, domain.PatronTypeRegular, []domain.Hold{
			newExampleHold(t),
		}, nil)

		err := patron.CancelHold(domain.BookID(uuid.NewString()))

		assert.ErrorIs(t, err, domain.ErrHoldNotFound)
	})

	t.Run("cancel a hold successfully", func(t *testing.T) {
		t.Parallel()
		hold := newExampleHold(t)
		patron := newExamplePatron(t, domain.PatronTypeRegular, []domain.Hold{hold}, nil)

		err := patron.CancelHold(hold.BookID)
		require.NoError(t, err)

		assert.Len(t, patron.Holds(), 0)
	})
}
