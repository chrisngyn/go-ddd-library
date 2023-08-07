package book

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrBookNotCheckedOut = commonErrors.NewIncorrectInputError("book-not-checked-out", "book not checked out")
)

// CheckIn checks in a book.
func (b *Book) CheckIn() error {
	if b.status != StatusCheckedOut {
		return ErrBookNotCheckedOut
	}

	b.checkedOutInfo = CheckedOutInformation{}
	b.status = StatusAvailable

	return nil
}
