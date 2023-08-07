package book_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestNewBookInformation_invalid(t *testing.T) {
	bookID := uuid.New()
	bookType := book.TypeCirculating
	placedAt := uuid.New()

	_, err := book.NewBookInformation(uuid.Nil, bookType, placedAt)
	assert.Error(t, err)

	_, err = book.NewBookInformation(bookID, book.Type{}, placedAt)
	assert.Error(t, err)

	_, err = book.NewBookInformation(bookID, bookType, uuid.Nil)
	assert.Error(t, err)
}
