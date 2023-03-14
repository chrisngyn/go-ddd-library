package query_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func TestOverdueCheckoutsHandler_Handle_invalid_query(t *testing.T) {
	h := query.OverdueCheckoutsHandler{}
	_, err := h.Handle(context.Background(), query.OverdueCheckoutsQuery{At: time.Time{}})
	assert.Error(t, err)
}
