package patron

import (
	"github.com/google/uuid"
)

func (p *Patron) Checkout(bookID uuid.UUID) error {
	return p.removeHold(bookID)
}
