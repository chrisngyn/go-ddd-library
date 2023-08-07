package patron_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func TestNewPatron_invalid(t *testing.T) {
	patronID := uuid.New()
	patronType := patron.TypeRegular

	_, err := patron.NewPatron(uuid.UUID{}, patronType, nil, nil)
	assert.Error(t, err)

	_, err = patron.NewPatron(patronID, patron.Type{}, nil, nil)
	assert.Error(t, err)
}
