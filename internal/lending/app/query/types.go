package query

import (
	"time"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type PatronProfile struct {
	PatronID         string
	PatronType       domain.PatronType
	Holds            []domain.Hold
	CheckedOuts      []CheckedOut
	OverdueCheckouts []OverdueCheckout
}

type CheckedOut struct {
	BookID          string
	LibraryBranchID string
	At              time.Time
}

type OverdueCheckout struct {
	PatronID        domain.PatronID
	BookID          string
	LibraryBranchID string
}

type ExpiredHold struct {
	BookID          domain.BookID
	LibraryBranchID domain.LibraryBranchID
	PatronID        domain.PatronID
	HoldTill        time.Time
}
