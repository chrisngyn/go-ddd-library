package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func TestCheckoutHandler_Handle_invalid_command(t *testing.T) {
	patronID := uuid.New()
	bookID := uuid.New()

	h := command.CheckoutHandler{}

	tests := []command.CheckoutCommand{
		{time.Time{}, patronID, bookID},
		{time.Now(), uuid.Nil, bookID},
		{time.Now(), patronID, uuid.Nil},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
