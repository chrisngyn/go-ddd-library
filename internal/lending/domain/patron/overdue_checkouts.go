package patron

import (
	"github.com/google/uuid"
)

// OverdueCheckouts is mapping of library branch id with list of book ids.
type OverdueCheckouts map[uuid.UUID][]uuid.UUID

func (o OverdueCheckouts) TotalAt(libraryBranchID uuid.UUID) int {
	return len(o[libraryBranchID])
}

func (o OverdueCheckouts) AddNewBookID(libraryBranchID uuid.UUID, bookID uuid.UUID) {
	o[libraryBranchID] = append(o[libraryBranchID], bookID)
}

func (o OverdueCheckouts) RemoveBookID(bookID uuid.UUID) {
	var foundLibraryBranchID uuid.UUID
	var bookIdx int
	var found bool
	for libraryBranchID, books := range o {
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

	overdueCheckouts := o[foundLibraryBranchID]
	overdueCheckouts = append(overdueCheckouts[:bookIdx], overdueCheckouts[bookIdx+1:]...)
	o[foundLibraryBranchID] = overdueCheckouts

	if len(overdueCheckouts) == 0 {
		delete(o, foundLibraryBranchID)
	}
}
