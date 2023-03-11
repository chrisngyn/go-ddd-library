package domain

type Hold struct {
	BookID       BookID
	PlacedAt     LibraryBranchID
	HoldDuration HoldDuration
}
