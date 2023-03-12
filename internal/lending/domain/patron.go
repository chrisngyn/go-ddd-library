package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

type PatronType struct {
	s string
}

func (t PatronType) IsZero() bool {
	return t == PatronType{}
}

var (
	PatronTypeRegular    = PatronType{"Regular"}
	PatronTypeResearcher = PatronType{"Researcher"}
)

type PatronID string

func (i PatronID) IsZero() bool {
	return i == ""
}

type Patron struct {
	id               PatronID
	patronType       PatronType
	holds            []Hold
	overdueCheckouts map[LibraryBranchID][]BookID
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

func NewPatron(id PatronID, patronType PatronType, holds []Hold, overdueCheckouts map[LibraryBranchID][]BookID) (Patron, error) {
	if id.IsZero() {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if patronType.IsZero() {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-type", "missing patron type")
	}
	return Patron{
		id:               id,
		patronType:       patronType,
		holds:            holds,
		overdueCheckouts: overdueCheckouts,
	}, nil
}
