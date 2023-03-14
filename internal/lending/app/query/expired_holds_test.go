package query_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func TestExpiredHoldsHandler_Handle_invalid_query(t *testing.T) {
	h := query.ExpiredHoldsHandler{}
	_, err := h.Handle(context.Background(), query.ExpiredHoldsQuery{At: time.Time{}})
	assert.Error(t, err)
}
