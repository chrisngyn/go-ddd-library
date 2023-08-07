package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestPlaceOnHoldHandler_Handle_invalid_command(t *testing.T) {
	patronID := uuid.New()
	bookID := uuid.New()
	holdDuration, err := patron.NewHoldDuration(time.Now(), 5)
	require.NoError(t, err)

	h := command.PlaceOnHoldHandler{}

	tests := []command.PlaceOnHoldCommand{
		{uuid.Nil, bookID, holdDuration},
		{patronID, uuid.Nil, holdDuration},
		{patronID, bookID, patron.HoldDuration{}},
	}

	for _, tt := range tests {
		err := h.Handle(context.Background(), tt)
		assert.Error(t, err)
	}
}
