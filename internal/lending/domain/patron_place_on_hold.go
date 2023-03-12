package domain

func (p *Patron) PlaceOnHold(book BookInformation, duration HoldDuration) error {
	if err := p.canHold(book, duration); err != nil {
		return err
	}

	p.holds = append(p.holds, Hold{
		BookID:       book.BookID,
		PlacedAt:     book.PlacedAt,
		HoldDuration: duration,
	})

	return nil
}

func (p *Patron) canHold(book BookInformation, duration HoldDuration) error {
	for _, policy := range holdPolices {
		if err := policy(book, p, duration); err != nil {
			return err
		}
	}
	return nil
}

func (p *Patron) isRegular() bool {
	return p.patronType == PatronTypeRegular
}

func (p *Patron) overdueCheckoutsAt(libraryBranchID LibraryBranchID) int {
	return len(p.overdueCheckouts[libraryBranchID])
}

func (p *Patron) numberOfHolds() int {
	return len(p.holds)
}
