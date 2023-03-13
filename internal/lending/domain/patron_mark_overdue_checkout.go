package domain

func (p *Patron) MarkOverdueCheckout(bookInfo BookInformation) {
	p.overdueCheckouts[bookInfo.PlacedAt] = append(p.overdueCheckouts[bookInfo.PlacedAt], bookInfo.BookID)
}
