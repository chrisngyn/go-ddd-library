package domain

import (
	"time"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotHoldByPatron = commonErrors.NewIncorrectInputError("book-not-hold-by-patron", "book not hold by patron")
)

func (b *Book) Checkout(patronID PatronID, at time.Time) error {
	if b.status != BookStatusOnHold {
		return ErrBookNotOnHold
	}

	if b.holdInfo.ByPatron != patronID {
		return errors.Wrapf(ErrBookNotHoldByPatron, "checkoutBy=%s holdBy=%s", patronID, b.holdInfo.ByPatron)
	}

	b.status = BookStatusCheckedOut
	b.holdInfo = HoldInformation{}
	b.checkedOutInfo = CheckedOutInformation{
		ByPatron: patronID,
		At:       at,
	}

	return nil
}
