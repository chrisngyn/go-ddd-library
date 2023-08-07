package patron_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestNewHold_invalid(t *testing.T) {
	bookID := uuid.New()
	libraryBranchID := uuid.New()
	holdDuration := newExampleCloseEndedHoldDuration(t)

	_, err := patron.NewHold(uuid.Nil, libraryBranchID, holdDuration)
	assert.Error(t, err)

	_, err = patron.NewHold(bookID, uuid.Nil, holdDuration)
	assert.Error(t, err)

	_, err = patron.NewHold(bookID, libraryBranchID, patron.HoldDuration{})
	assert.Error(t, err)
}
