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

var (
	BookTypeRestricted  = BookType{"Restricted"}
	BookTypeCirculating = BookType{"Circulating"}
)

type Book struct {
	bookInformation           BookInformation
	status                    BookStatus
	bookHoldInformation       BookHoldInformation
	BookCheckedOutInformation BookCheckedOutInformation
}

func (b *Book) BookInformation() BookInformation {
	return b.bookInformation
}

type BookInformation struct {
	BookID   BookID
	BookType BookType
	PlacedAt LibraryBranchID
}

func (b BookInformation) IsRestricted() bool {
	return b.BookType == BookTypeRestricted
}

type BookHoldInformation struct {
	ByPatron PatronID
	Till     time.Time
}

type BookCheckedOutInformation struct {
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

	b.bookHoldInformation = BookHoldInformation{
		ByPatron: patronID,
		Till:     holdDuration.till,
	}
	b.status = BookStatusOnHold
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

	if b.bookHoldInformation.ByPatron != patronID {
		return errors.Wrapf(ErrBookNotHoldByPatron, "checkoutBy=%s holdBy=%s", patronID, b.bookHoldInformation.ByPatron)
	}

	b.status = BookStatusCheckedOut
	b.BookCheckedOutInformation = BookCheckedOutInformation{
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

	b.bookHoldInformation = BookHoldInformation{}
	b.BookCheckedOutInformation = BookCheckedOutInformation{}
	b.status = BookStatusAvailable

	return nil
}
