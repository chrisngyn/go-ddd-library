package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestNewHold_invalid(t *testing.T) {
	bookID := domain.BookID(uuid.NewString())
	libraryBranchID := domain.LibraryBranchID(uuid.NewString())
	holdDuration := newExampleCloseEndedHoldDuration(t)

	_, err := domain.NewHold("", libraryBranchID, holdDuration)
	assert.Error(t, err)

	_, err = domain.NewHold(bookID, "", holdDuration)
	assert.Error(t, err)

	_, err = domain.NewHold(bookID, libraryBranchID, domain.HoldDuration{})
	assert.Error(t, err)
}
