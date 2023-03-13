package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotAvailable = commonErrors.NewIncorrectInputError("book-not-available", "book not available")
)

func (b *Book) HoldBy(patronID PatronID, holdDuration HoldDuration) error {
	if b.status != BookStatusAvailable {
		return ErrBookNotAvailable
	}

	b.holdInfo = HoldInformation{
		ByPatron: patronID,
		Till:     holdDuration.till,
	}
	b.status = BookStatusOnHold
	return nil
}
