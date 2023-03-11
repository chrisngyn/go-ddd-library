package domain

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrHoldNotFound = commonErrors.NewIncorrectInputError("hold-not-found", "hold not found")
)

func (p *Patron) CancelHold(bookID BookID) error {
	return p.removeHold(bookID)
}

func (p *Patron) removeHold(bookID BookID) error {
	var idx int
	var found bool
	for i, h := range p.holds {
		if h.BookID == bookID {
			idx = i
			found = true
			break
		}
	}

	if !found {
		return ErrHoldNotFound
	}

	p.holds = append(p.holds[:idx], p.holds[idx+1:]...)
	return nil
}
