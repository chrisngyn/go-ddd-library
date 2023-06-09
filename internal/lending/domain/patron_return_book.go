package domain

func (p *Patron) ReturnBook(bookID BookID) {
	p.removeOverdueCheckoutIfExist(bookID)
}

func (p *Patron) removeOverdueCheckoutIfExist(bookID BookID) {
	var foundLibraryBranchID LibraryBranchID
	var bookIdx int
	var found bool
	for libraryBranchID, books := range p.overdueCheckouts {
		for i, b := range books {
			if bookID == b {
				foundLibraryBranchID = libraryBranchID
				bookIdx = i
				found = true
				break
			}
		}
	}

	if !found {
		return
	}

	overdueCheckouts := p.overdueCheckouts[foundLibraryBranchID]
	overdueCheckouts = append(overdueCheckouts[:bookIdx], overdueCheckouts[bookIdx+1:]...)

	if len(overdueCheckouts) == 0 {
		delete(p.overdueCheckouts, foundLibraryBranchID)
	}

	p.overdueCheckouts[foundLibraryBranchID] = overdueCheckouts
}
