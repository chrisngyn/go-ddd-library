package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func TestMarkOverdueCheckoutHandler_Handle_invalid_command(t *testing.T) {
	patronID := uuid.New()
	bookID := uuid.New()
	libraryBranchID := uuid.New()

	h := command.MarkOverdueCheckoutHandler{}

	tests := []command.MarkOverdueCheckoutCommand{
		{uuid.Nil, bookID, libraryBranchID},
		{patronID, uuid.Nil, libraryBranchID},
		{patronID, bookID, uuid.Nil},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
