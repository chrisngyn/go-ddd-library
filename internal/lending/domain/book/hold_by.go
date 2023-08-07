package book

import (
	"time"

	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotAvailable = commonErrors.NewIncorrectInputError("book-not-available", "book not available")
)

// HoldBy mark a book is hold by a patron.
func (b *Book) HoldBy(patronID uuid.UUID, till time.Time) error {
	if b.status != StatusAvailable {
		return ErrBookNotAvailable
	}

	b.holdInfo = HoldInformation{
		ByPatron: patronID,
		Till:     till,
	}
	b.status = StatusOnHold
	return nil
}
