package book

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotHoldByPatron = commonErrors.NewIncorrectInputError("book-not-hold-by-patron", "book not hold by patron")
)

// Checkout checks out a book.
func (b *Book) Checkout(patronID uuid.UUID, at time.Time) error {
	if b.status != StatusOnHold {
		return ErrBookNotOnHold
	}

	if b.holdInfo.ByPatron != patronID {
		return errors.Wrapf(ErrBookNotHoldByPatron, "checkoutBy=%s holdBy=%s", patronID, b.holdInfo.ByPatron)
	}

	b.status = StatusCheckedOut
	b.holdInfo = HoldInformation{}
	b.checkedOutInfo = CheckedOutInformation{
		ByPatron: patronID,
		At:       at,
	}

	return nil
}
