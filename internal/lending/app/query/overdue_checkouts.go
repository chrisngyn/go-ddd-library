package query

import (
	"context"
	"time"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

const (
	maxCheckoutDurationDays = 60
)

type OverdueCheckoutsHandler struct {
	readModel OverdueCheckoutsReadModel
}

func NewOverdueCheckoutsHandler(readModel OverdueCheckoutsReadModel) OverdueCheckoutsHandler {
	if readModel == nil {
		panic("missing readModel")
	}
	return OverdueCheckoutsHandler{readModel: readModel}
}

type OverdueCheckoutsReadModel interface {
	ListOverdueCheckouts(ctx context.Context, at time.Time, maxCheckoutDurationDays int) ([]OverdueCheckout, error)
}

func (h OverdueCheckoutsHandler) Handle(ctx context.Context, query OverdueCheckoutsQuery) ([]OverdueCheckout, error) {
	if err := query.validate(); err != nil {
		return nil, errors.Wrap(err, "validate")
	}
	return h.readModel.ListOverdueCheckouts(ctx, query.At, maxCheckoutDurationDays)
}

type OverdueCheckoutsQuery struct {
	At time.Time
}

func (q OverdueCheckoutsQuery) validate() error {
	if q.At.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-at", "missing at")
	}
	return nil
}
