package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestNewBookInformation_invalid(t *testing.T) {
	bookID := domain.BookID(uuid.NewString())
	bookType := domain.BookTypeCirculating
	placedAt := domain.LibraryBranchID(uuid.NewString())

	_, err := domain.NewBookInformation("", bookType, placedAt)
	assert.Error(t, err)

	_, err = domain.NewBookInformation(bookID, domain.BookType{}, placedAt)
	assert.Error(t, err)

	_, err = domain.NewBookInformation(bookID, bookType, "")
	assert.Error(t, err)
}
