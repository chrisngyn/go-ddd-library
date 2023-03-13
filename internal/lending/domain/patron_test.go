package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestNewPatron_invalid(t *testing.T) {
	patronID := domain.PatronID("regular1")
	patronType := domain.PatronTypeRegular

	_, err := domain.NewPatron("", patronType, nil, nil)
	assert.Error(t, err)

	_, err = domain.NewPatron(patronID, domain.PatronType{}, nil, nil)
	assert.Error(t, err)
}
