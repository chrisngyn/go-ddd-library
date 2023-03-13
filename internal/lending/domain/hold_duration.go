package domain

import (
	"time"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type HoldDuration struct {
	from time.Time
	till time.Time
}

func (d HoldDuration) From() time.Time {
	return d.from
}

func (d HoldDuration) Till() time.Time {
	return d.till
}

func (d HoldDuration) IsZero() bool {
	return d == HoldDuration{}
}

func NewHoldDuration(from time.Time, numOfDays int) (HoldDuration, error) {
	if from.IsZero() {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("missing-from", "missing from")
	}
	if numOfDays < 0 {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("invalid-num-of-days", "numOfDays must great than 0")
	}

	var till time.Time
	if numOfDays > 0 {
		till = from.AddDate(0, 0, numOfDays)
	}
	return NewHoldDurationFromTill(from, till)
}

func NewHoldDurationFromTill(from time.Time, till time.Time) (HoldDuration, error) {
	if from.IsZero() {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("missing-from", "missing from")
	}
	if !till.IsZero() && till.Before(from) {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("invalid-till", "till must after from")
	}
	return HoldDuration{
		from: from,
		till: till,
	}, nil
}

func (d HoldDuration) IsOpenEnded() bool {
	return d.till.IsZero()
}
