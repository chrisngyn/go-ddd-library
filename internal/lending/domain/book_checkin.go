package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotCheckedOut = commonErrors.NewIncorrectInputError("book-not-checked-out", "book not checked out")
)

func (b *Book) CheckIn() error {
	if b.status != BookStatusCheckedOut {
		return ErrBookNotCheckedOut
	}

	b.checkedOutInfo = CheckedOutInformation{}
	b.status = BookStatusAvailable

	return nil
}
