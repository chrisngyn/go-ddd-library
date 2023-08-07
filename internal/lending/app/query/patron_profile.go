package query

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/common/monitoring"
)

type PatronProfileHandler struct {
	readModel PatronProfileReadModel
}

func NewPatronProfileHandler(readModel PatronProfileReadModel) PatronProfileHandler {
	if readModel == nil {
		panic("missing real model")
	}
	return PatronProfileHandler{readModel: readModel}
}

type PatronProfileReadModel interface {
	GetPatronProfile(ctx context.Context, patronID uuid.UUID) (PatronProfile, error)
}

func (h PatronProfileHandler) Handle(ctx context.Context, query PatronProfileQuery) (p PatronProfile, err error) {
	defer func(st time.Time) {
		monitoring.MonitorQuery(ctx, "PatronProfile", query, p, err, st)
	}(time.Now())

	if err := query.validate(); err != nil {
		return p, errors.Wrap(err, "validate query")
	}

	return h.readModel.GetPatronProfile(ctx, query.PatronID)
}

type PatronProfileQuery struct {
	PatronID uuid.UUID
}

func (q PatronProfileQuery) validate() error {
	if q.PatronID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	return nil
}
