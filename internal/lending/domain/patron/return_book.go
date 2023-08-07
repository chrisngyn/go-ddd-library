package patron

import (
	"github.com/google/uuid"
)

func (p *Patron) ReturnBook(bookID uuid.UUID) {
	p.overdueCheckouts.RemoveBookID(bookID)
}
