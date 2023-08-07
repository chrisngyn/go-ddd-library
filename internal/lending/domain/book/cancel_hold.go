package book

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotOnHold = commonErrors.NewIncorrectInputError("book-not-on-hold", "book not on hold")
)

// CancelHold cancels a book on hold.
func (b *Book) CancelHold() error {
	if b.status != StatusOnHold {
		return ErrBookNotOnHold
	}

	b.holdInfo = HoldInformation{}
	b.status = StatusAvailable

	return nil
}
