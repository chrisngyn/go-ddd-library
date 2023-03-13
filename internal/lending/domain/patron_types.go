package domain

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
