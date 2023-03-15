package query

import (
	"context"
	"time"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/common/monitoring"
)

type ExpiredHoldsHandler struct {
	readModel ExpiredHoldsReadModel
}

func NewExpiredHoldsHandler(readModel ExpiredHoldsReadModel) ExpiredHoldsHandler {
	if readModel == nil {
		panic("missing readModel")
	}
	return ExpiredHoldsHandler{readModel: readModel}
}

type ExpiredHoldsReadModel interface {
	ListExpiredHolds(ctx context.Context, at time.Time) ([]ExpiredHold, error)
}

func (h ExpiredHoldsHandler) Handle(ctx context.Context, query ExpiredHoldsQuery) (result []ExpiredHold, err error) {
	defer func(st time.Time) {
		monitoring.MonitorQuery(ctx, "ExpiredHolds", query, result, err, st)
	}(time.Now())

	if err := query.validate(); err != nil {
		return nil, errors.Wrap(err, "validate")
	}

	return h.readModel.ListExpiredHolds(ctx, query.At)
}

type ExpiredHoldsQuery struct {
	At time.Time
}

func (q ExpiredHoldsQuery) validate() error {
	if q.At.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-at", "missing at")
	}
	return nil
}
