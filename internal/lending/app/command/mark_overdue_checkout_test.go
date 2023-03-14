package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestMarkOverdueCheckoutHandler_Handle_invalid_command(t *testing.T) {
	patronID := domain.PatronID(uuid.NewString())
	bookID := domain.BookID(uuid.NewString())
	libraryBranchID := domain.LibraryBranchID(uuid.NewString())

	h := command.MarkOverdueCheckoutHandler{}

	tests := []command.MarkOverdueCheckoutCommand{
		{"", bookID, libraryBranchID},
		{patronID, "", libraryBranchID},
		{patronID, bookID, ""},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
