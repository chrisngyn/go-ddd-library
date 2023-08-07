package patron

import (
	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type Hold struct {
	BookID       uuid.UUID
	PlacedAt     uuid.UUID
	HoldDuration HoldDuration
}

func NewHold(bookID uuid.UUID, placedAt uuid.UUID, duration HoldDuration) (Hold, error) {
	if bookID == uuid.Nil {
		return Hold{}, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if placedAt == uuid.Nil {
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
