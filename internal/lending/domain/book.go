package domain

import (
	"time"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type BookID string

func (i BookID) IsZero() bool {
	return i == ""
}

type LibraryBranchID string

func (i LibraryBranchID) IsZero() bool {
	return i == ""
}

type BookType struct {
	s string
}

func (t BookType) IsZero() bool {
	return t == BookType{}
}

var (
	BookTypeRestricted  = BookType{"Restricted"}
	BookTypeCirculating = BookType{"Circulating"}
)

type Book struct {
	info   Information
	status BookStatus

	holdInfo       HoldInformation
	checkedOutInfo CheckedOutInformation
}

func (b *Book) ID() BookID {
	return b.info.BookID
}

func (b *Book) BookInfo() Information {
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

func NewAvailableBook(information Information) (Book, error) {
	if information == (Information{}) {
		return Book{}, commonErrors.NewIncorrectInputError("missing-information", "missing book information")
	}
	return Book{
		info:   information,
		status: BookStatusAvailable,
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
		status:   BookStatusOnHold,
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
		status:         BookStatusCheckedOut,
		checkedOutInfo: checkedOutInformation,
	}, nil
}

type Information struct {
	BookID   BookID
	BookType BookType
	PlacedAt LibraryBranchID
}

func NewBookInformation(bookID BookID, bookType BookType, placedAt LibraryBranchID) (Information, error) {
	if bookID.IsZero() {
		return Information{}, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if bookType.IsZero() {
		return Information{}, commonErrors.NewIncorrectInputError("missing-book-type", "missing book type")
	}
	if placedAt.IsZero() {
		return Information{}, commonErrors.NewIncorrectInputError("missing-placed-at", "missing placed at")
	}
	return Information{
		BookID:   bookID,
		BookType: bookType,
		PlacedAt: placedAt,
	}, nil
}

func (b Information) IsRestricted() bool {
	return b.BookType == BookTypeRestricted
}

type HoldInformation struct {
	ByPatron PatronID
	Till     time.Time
}

type CheckedOutInformation struct {
	ByPatron PatronID
	At       time.Time
}

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

func (b *Book) CancelHold() error {
	if b.status != BookStatusOnHold {
		return ErrBookNotOnHold
	}

	b.holdInfo = HoldInformation{}
	b.status = BookStatusAvailable

	return nil
}

var (
	ErrBookNotOnHold       = commonErrors.NewIncorrectInputError("book-not-on-hold", "book not on hold")
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
