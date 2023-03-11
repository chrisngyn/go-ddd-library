package domain

func (p *Patron) Checkout(bookID BookID) error {
	return p.removeHold(bookID)
}
