package query

import (
	"time"

	"github.com/google/uuid"

	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

// Those structs here are read models. It used as for mapping data from many domain entities and return client.
// There are good reason to separate it with domain model because it's usually not contain complex business logic.
// And easy change depends on client requirement.

type PatronProfile struct {
	PatronID         uuid.UUID
	PatronType       patron.Type
	Holds            []patron.Hold
	CheckedOuts      []CheckedOut
	OverdueCheckouts []OverdueCheckout
}

type CheckedOut struct {
	BookID          uuid.UUID
	LibraryBranchID uuid.UUID
	At              time.Time
}

type OverdueCheckout struct {
	PatronID        uuid.UUID
	BookID          uuid.UUID
	LibraryBranchID uuid.UUID
}

type ExpiredHold struct {
	BookID          uuid.UUID
	LibraryBranchID uuid.UUID
	PatronID        uuid.UUID
	HoldTill        time.Time
}
