package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type Hold struct {
	BookID       BookID
	PlacedAt     LibraryBranchID
	HoldDuration HoldDuration
}

func NewHold(bookID BookID, placedAt LibraryBranchID, duration HoldDuration) (Hold, error) {
	if bookID.IsZero() {
		return Hold{}, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if placedAt.IsZero() {
		return Hold{}, commonErrors.NewIncorrectInputError("missing-placed-at", "missing-placed-at")
	}
	if duration.IsZero() {
		return Hold{}, commonErrors.NewIncorrectInputError("missing-hold-duration", "missing hold duration")
	}
	return Hold{
		BookID:       bookID,
		PlacedAt:     placedAt,
		HoldDuration: duration,
	}, nil
}
