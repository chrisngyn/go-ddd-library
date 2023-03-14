package query_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func TestPatronProfileHandler_Handle_invalid_query(t *testing.T) {
	h := query.PatronProfileHandler{}
	_, err := h.Handle(context.Background(), query.PatronProfileQuery{PatronID: ""})
	assert.Error(t, err)
}
