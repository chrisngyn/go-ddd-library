package patron

import (
	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

// Patron is a person who borrows books from the library. It is an aggregate root.
// And it contains some information to determine if a patron can hold a book or not.
type Patron struct {
	id               uuid.UUID
	patronType       Type
	holds            []Hold
	overdueCheckouts OverdueCheckouts
}

func NewPatron(id uuid.UUID, patronType Type, holds []Hold, overdueCheckouts OverdueCheckouts) (Patron, error) {
	if id == uuid.Nil {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if patronType.IsZero() {
		return Patron{}, commonErrors.NewIncorrectInputError("missing-patron-type", "missing patron type")
	}

	if overdueCheckouts == nil {
		overdueCheckouts = make(map[uuid.UUID][]uuid.UUID)
	}
	return Patron{
		id:               id,
		patronType:       patronType,
		holds:            holds,
		overdueCheckouts: overdueCheckouts,
	}, nil
}

func (p *Patron) ID() uuid.UUID {
	return p.id
}

func (p *Patron) PatronType() Type {
	return p.patronType
}

func (p *Patron) Holds() []Hold {
	return p.holds
}

func (p *Patron) OverdueCheckouts() OverdueCheckouts {
	return p.overdueCheckouts
}

func (p *Patron) IsZero() bool {
	return p.id == uuid.Nil
}

func (p *Patron) isRegular() bool {
	return p.patronType == TypeRegular
}

func (p *Patron) overdueCheckoutsAt(libraryBranchID uuid.UUID) int {
	return p.overdueCheckouts.TotalAt(libraryBranchID)
}

func (p *Patron) numberOfHolds() int {
	return len(p.holds)
}
