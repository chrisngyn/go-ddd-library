package domain

type BookStatus struct {
	s string
}

var (
	BookStatusAvailable  = BookStatus{"Available"}
	BookStatusOnHold     = BookStatus{"OnHold"}
	BookStatusCheckedOut = BookStatus{"CheckedOut"}
)
