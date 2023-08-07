package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestAddNewBookHandler_Handle_invalid_command(t *testing.T) {
	bookID := uuid.New()
	bookType := book.TypeRestricted
	libraryBranchID := uuid.New()

	h := command.AddNewBookHandler{}

	tests := []command.AddNewBookCommand{
		{uuid.Nil, bookType, libraryBranchID},
		{bookID, book.Type{}, libraryBranchID},
		{bookID, bookType, uuid.Nil},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
