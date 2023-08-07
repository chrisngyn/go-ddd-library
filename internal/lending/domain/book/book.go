package book

import (
	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

// Book is the aggregate root of the Book domain.
// It contains book information, and other information of the book when it is hold or checked out.
// In the initial idea, I want to have a separate aggregate for each book status, but I think it is not necessary.
// And it will introduce more duplicate code. Example, when an AvailableBook is hold, we need to store hold information in it.
// And it seems duplicate with the BookOnHold aggregate.
type Book struct {
	info   Information
	status Status

	holdInfo       HoldInformation
	checkedOutInfo CheckedOutInformation
}

func NewAvailableBook(information Information) (Book, error) {
	if information == (Information{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	return Book{
		info:   information,
		status: StatusAvailable,
	}, nil
}

func NewBookOnHold(information Information, holdInformation HoldInformation) (Book, error) {
	if information == (Information{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	if holdInformation == (HoldInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-hold-information", "missing hold information")
	}
	return Book{
		info:     information,
		status:   StatusOnHold,
		holdInfo: holdInformation,
	}, nil
}

func NewCheckedOutBook(information Information, checkedOutInformation CheckedOutInformation) (Book, error) {
	if information == (Information{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	if checkedOutInformation == (CheckedOutInformation{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-checked-out-information", "missing checked out information")
	}
	return Book{
		info:           information,
		status:         StatusCheckedOut,
		checkedOutInfo: checkedOutInformation,
	}, nil
}

func (b *Book) ID() uuid.UUID {
	return b.info.BookID
}

func (b *Book) BookInfo() Information {
	return b.info
}

func (b *Book) Status() Status {
	return b.status
}

func (b *Book) ByPatronID() uuid.UUID {
	if b.status == StatusOnHold {
		return b.holdInfo.ByPatron
	}
	if b.status == StatusCheckedOut {
		return b.checkedOutInfo.ByPatron
	}
	return uuid.Nil
}

func (b *Book) BookHoldInfo() HoldInformation {
	return b.holdInfo
}

func (b *Book) BookCheckedOutInfo() CheckedOutInformation {
	return b.checkedOutInfo
}
