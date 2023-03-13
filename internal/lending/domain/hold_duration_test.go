package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestNewHoldDuration(t *testing.T) {
	t.Run("close-ended hold duration", func(t *testing.T) {
		t.Parallel()
		startTime := time.Now()
		numOfDays := 10

		holdDuration, err := domain.NewHoldDuration(startTime, numOfDays)
		require.NoError(t, err)

		assert.False(t, holdDuration.IsZero())
		assert.False(t, holdDuration.IsOpenEnded())
		assert.Equal(t, startTime, holdDuration.From())
		assert.Equal(t, startTime.AddDate(0, 0, numOfDays), holdDuration.Till())
	})
	t.Run("open-ended hold duration", func(t *testing.T) {
		t.Parallel()
		startTime := time.Now()

		holdDuration, err := domain.NewHoldDuration(startTime, 0)
		require.NoError(t, err)

		assert.False(t, holdDuration.IsZero())
		assert.True(t, holdDuration.IsOpenEnded())
		assert.Equal(t, startTime, holdDuration.From())
		assert.True(t, holdDuration.Till().IsZero())
	})
}

func TestNewHoldDuration_invalid(t *testing.T) {
	t.Parallel()
	from := time.Now()
	numOfDays := 10

	_, err := domain.NewHoldDuration(time.Time{}, numOfDays)
	assert.Error(t, err)

	_, err = domain.NewHoldDuration(from, -1)
	assert.Error(t, err)
}

func TestNewHoldDurationFromTill_invalid(t *testing.T) {
	t.Parallel()
	from := time.Now()

	_, err := domain.NewHoldDurationFromTill(time.Time{}, from.AddDate(0, 0, 1))
	assert.Error(t, err)

	_, err = domain.NewHoldDurationFromTill(from, from.AddDate(0, 0, -1))
	assert.Error(t, err)
}
