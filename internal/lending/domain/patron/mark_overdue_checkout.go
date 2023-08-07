package patron

import (
	"github.com/google/uuid"
)

func (p *Patron) MarkOverdueCheckout(bookID uuid.UUID, libraryBranchID uuid.UUID) {
	p.overdueCheckouts.AddNewBookID(libraryBranchID, bookID)
}
