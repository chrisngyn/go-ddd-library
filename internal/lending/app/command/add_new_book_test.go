package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestAddNewBookHandler_Handle_invalid_command(t *testing.T) {
	bookID := domain.BookID(uuid.NewString())
	bookType := domain.BookTypeRestricted
	libraryBranchID := domain.LibraryBranchID(uuid.NewString())

	h := command.AddNewBookHandler{}

	tests := []command.AddNewBookCommand{
		{"", bookType, libraryBranchID},
		{bookID, domain.BookType{}, libraryBranchID},
		{bookID, bookType, ""},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
