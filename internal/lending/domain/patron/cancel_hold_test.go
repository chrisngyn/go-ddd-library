package patron_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPatron_CancelHold(t *testing.T) {
	t.Run("cancel a hold that not existed", func(t *testing.T) {
		t.Parallel()
		aPatron := newExamplePatron(t, patron.TypeRegular, []patron.Hold{
			newExampleHold(t),
		}, nil)

		err := aPatron.CancelHold(uuid.New())

		assert.ErrorIs(t, err, patron.ErrHoldNotFound)
	})

	t.Run("cancel a hold successfully", func(t *testing.T) {
		t.Parallel()
		hold := newExampleHold(t)
		aPatron := newExamplePatron(t, patron.TypeRegular, []patron.Hold{hold}, nil)

		err := aPatron.CancelHold(hold.BookID)
		require.NoError(t, err)

		assert.Len(t, aPatron.Holds(), 0)
	})
}
