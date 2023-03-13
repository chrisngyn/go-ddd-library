package domain

import (
	"time"

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

type BookStatus struct {
	s string
}

var (
	BookStatusAvailable  = BookStatus{"Available"}
	BookStatusOnHold     = BookStatus{"OnHold"}
	BookStatusCheckedOut = BookStatus{"CheckedOut"}
)

type BookInformation struct {
	BookID   BookID
	BookType BookType
	PlacedAt LibraryBranchID
}

func NewBookInformation(bookID BookID, bookType BookType, placedAt LibraryBranchID) (BookInformation, error) {
	if bookID.IsZero() {
		return BookInformation{}, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if bookType.IsZero() {
		return BookInformation{}, commonErrors.NewIncorrectInputError("missing-book-type", "missing book type")
	}
	if placedAt.IsZero() {
		return BookInformation{}, commonErrors.NewIncorrectInputError("missing-placed-at", "missing placed at")
	}
	return BookInformation{
		BookID:   bookID,
		BookType: bookType,
		PlacedAt: placedAt,
	}, nil
}

func (b BookInformation) IsRestricted() bool {
	return b.BookType == BookTypeRestricted
}

type HoldInformation struct {
	ByPatron PatronID
	Till     time.Time
}

func (h HoldInformation) IsZero() bool {
	return h == HoldInformation{}
}

type CheckedOutInformation struct {
	ByPatron PatronID
	At       time.Time
}

func (c CheckedOutInformation) IsZero() bool {
	return c == CheckedOutInformation{}
}
