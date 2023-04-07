package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

// Patron is a person who borrows books from the library. It is an aggregate root.
// And it contains some information to determine if a patron can hold a book or not.
type Patron struct {
	id               PatronID
	patronType       PatronType
	holds            []Hold
	overdueCheckouts map[LibraryBranchID][]BookID
}

func NewPatron(id PatronID, patronType PatronType, holds []Hold, overdueCheckouts map[LibraryBranchID][]BookID) (Patron, error) {
	if id.IsZero() {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if patronType.IsZero() {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-type", "missing patron type")
	}

	if overdueCheckouts == nil {
		overdueCheckouts = make(map[LibraryBranchID][]BookID)
	}
	return Patron{
		id:               id,
		patronType:       patronType,
		holds:            holds,
		overdueCheckouts: overdueCheckouts,
	}, nil
}

func (p *Patron) ID() PatronID {
	return p.id
}

func (p *Patron) PatronType() PatronType {
	return p.patronType
}

func (p *Patron) Holds() []Hold {
	return p.holds
}

func (p *Patron) OverdueCheckouts() map[LibraryBranchID][]BookID {
	return p.overdueCheckouts
}

func (p *Patron) IsZero() bool {
	return p.id.IsZero()
}

func (p *Patron) isRegular() bool {
	return p.patronType == PatronTypeRegular
}

func (p *Patron) overdueCheckoutsAt(libraryBranchID LibraryBranchID) int {
	return len(p.overdueCheckouts[libraryBranchID])
}

func (p *Patron) numberOfHolds() int {
	return len(p.holds)
}
