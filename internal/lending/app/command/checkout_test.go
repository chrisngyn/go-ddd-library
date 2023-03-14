package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestCheckoutHandler_Handle_invalid_command(t *testing.T) {
	patronID := domain.PatronID(uuid.NewString())
	bookID := domain.BookID(uuid.NewString())

	h := command.CheckoutHandler{}

	tests := []command.CheckoutCommand{
		{time.Time{}, patronID, bookID},
		{time.Now(), "", bookID},
		{time.Now(), patronID, ""},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
