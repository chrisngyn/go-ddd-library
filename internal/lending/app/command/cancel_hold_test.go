package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func TestCancelHoldHandler_Handle_invalid_command(t *testing.T) {
	patronID := uuid.New()
	bookID := uuid.New()

	h := command.CancelHoldHandler{}

	tests := []command.CancelHoldCommand{
		{uuid.Nil, bookID},
		{patronID, uuid.Nil},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
