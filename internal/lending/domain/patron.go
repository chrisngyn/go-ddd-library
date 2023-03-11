package domain

type PatronType struct {
	s string
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
	overdueCheckouts map[LibraryBranchID][]BookID
	holds            []Hold
}

func (p *Patron) ID() PatronID {
	return p.id
}

func (p *Patron) IsZero() bool {
	return p.id.IsZero()
}
