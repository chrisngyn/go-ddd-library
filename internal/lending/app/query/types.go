package query

import (
	"time"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

// Those structs here are read models. It used as for mapping data from many domain entities and return client.
// There are good reason to separate it with domain model because it's usually not contain complex business logic.
// And easy change depends on client requirement.

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
