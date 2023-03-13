package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotOnHold = commonErrors.NewIncorrectInputError("book-not-on-hold", "book not on hold")
)

func (b *Book) CancelHold() error {
	if b.status != BookStatusOnHold {
		return ErrBookNotOnHold
	}

	b.holdInfo = HoldInformation{}
	b.status = BookStatusAvailable

	return nil
}
