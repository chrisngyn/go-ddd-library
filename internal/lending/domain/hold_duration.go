package domain

import (
	"time"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type HoldDuration struct {
	from time.Time
	till time.Time
}

func (d HoldDuration) IsZero() bool {
	return d == HoldDuration{}
}

func NewHoldDuration(from time.Time, numOfDays int) (HoldDuration, error) {
	if from.IsZero() {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("missing-from", "missing from")
	}
	if numOfDays < 0 {
		return HoldDuration{}, commonErrors.NewIncorrectInputError("invalid-num-of-days", "invalid-num-of-days")
	}

	holdDuration := HoldDuration{from: from}
	if numOfDays > 0 {
		holdDuration.till = from.AddDate(0, 0, numOfDays)
	}
	return holdDuration, nil
}

func (d HoldDuration) IsOpenEnded() bool {
	return d.till.IsZero()
}
