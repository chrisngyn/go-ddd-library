package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

// Book is the aggregate root of the Book domain.
// It contains book information, and other information of the book when it is hold or checked out.
// In the initial idea, I want to have a separate aggregate for each book status, but I think it is not necessary.
// And it will introduce more duplicate code. Example, when an AvailableBook is hold, we need to store hold information in it.
// And it seems duplicate with the BookOnHold aggregate.
type Book struct {
	info   BookInformation
	status BookStatus

	holdInfo       HoldInformation
	checkedOutInfo CheckedOutInformation
}

func NewAvailableBook(information BookInformation) (Book, error) {
	if information == (BookInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	return Book{
		info:   information,
		status: BookStatusAvailable,
	}, nil
}

func NewBookOnHold(information BookInformation, holdInformation HoldInformation) (Book, error) {
	if information == (BookInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	if holdInformation == (HoldInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-hold-information", "missing hold information")
	}
	return Book{
		info:     information,
		status:   BookStatusOnHold,
		holdInfo: holdInformation,
	}, nil
}

func NewCheckedOutBook(information BookInformation, checkedOutInformation CheckedOutInformation) (Book, error) {
	if information == (BookInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	if checkedOutInformation == (CheckedOutInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-checked-out-information", "missing checked out information")
	}
	return Book{
		info:           information,
		status:         BookStatusCheckedOut,
		checkedOutInfo: checkedOutInformation,
	}, nil
}

func (b *Book) ID() BookID {
	return b.info.BookID
}

func (b *Book) BookInfo() BookInformation {
	return b.info
}

func (b *Book) Status() BookStatus {
	return b.status
}

func (b *Book) ByPatronID() PatronID {
	if b.status == BookStatusOnHold {
		return b.holdInfo.ByPatron
	}
	if b.status == BookStatusCheckedOut {
		return b.checkedOutInfo.ByPatron
	}
	return ""
}

func (b *Book) BookHoldInfo() HoldInformation {
	return b.holdInfo
}

func (b *Book) BookCheckedOutInfo() CheckedOutInformation {
	return b.checkedOutInfo
}
