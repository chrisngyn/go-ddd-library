package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func TestReturnBookHandler_Handle_invalid_command(t *testing.T) {
	h := command.ReturnBookHandler{}
	err := h.Handle(context.Background(), command.ReturnBookCommand{BookID: uuid.Nil})
	assert.Error(t, err)
}
