package patron

type Type struct {
	s string
}

func (t Type) IsZero() bool {
	return t == Type{}
}

var (
	TypeRegular    = Type{"Regular"}
	TypeResearcher = Type{"Researcher"}
)
