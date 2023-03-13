package domain

func (p *Patron) MarkOverdueCheckout(bookID BookID, libraryBranchID LibraryBranchID) {
	p.overdueCheckouts[libraryBranchID] = append(p.overdueCheckouts[libraryBranchID], bookID)
}
