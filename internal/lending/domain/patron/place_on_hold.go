package patron

import (
	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func (p *Patron) PlaceOnHold(book book.Information, duration HoldDuration) error {
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

func (p *Patron) canHold(book book.Information, duration HoldDuration) error {
	for _, policy := range holdPolices {
		if err := policy(book, p, duration); err != nil {
			return err
		}
	}
	return nil
}

type HoldPolicy func(book book.Information, patron *Patron, duration HoldDuration) error

var (
	holdPolices = []HoldPolicy{
		onlyResearcherPatronsCanHoldRestrictedBooksPolicy,
		overdueCheckoutsRejectionPolicy,
		regularPatronMaximumNumberOfHoldsPolicy,
		onlyResearcherPatronsCanPlaceOpenEndedHoldsPolicy,
	}
)

var ErrRegularPatronCannotHoldRestrictedBook = commonErrors.NewIncorrectInputError(
	"regular-patron-cannot-hold-restricted-book",
	"regular patron cannot hold restricted book",
)

func onlyResearcherPatronsCanHoldRestrictedBooksPolicy(book book.Information, patron *Patron, _ HoldDuration) error {
	if book.IsRestricted() && patron.isRegular() {
		return ErrRegularPatronCannotHoldRestrictedBook
	}
	return nil
}

const maxCountOfOverdueCheckouts = 2

var ErrMaxCountOfOverdueCheckoutsReached = commonErrors.NewIncorrectInputError(
	"max-count-of-overdue-checkouts-reached",
	"max count of overdue checkout reached",
)

func overdueCheckoutsRejectionPolicy(book book.Information, patron *Patron, _ HoldDuration) error {
	if patron.overdueCheckoutsAt(book.PlacedAt) >= maxCountOfOverdueCheckouts {
		return ErrMaxCountOfOverdueCheckoutsReached
	}
	return nil
}

const maxNumberOfHolds = 5

var ErrMaxHoldsReached = commonErrors.NewIncorrectInputError(
	"max-holds-reached",
	"patron cannot hold more books",
)

func regularPatronMaximumNumberOfHoldsPolicy(_ book.Information, patron *Patron, _ HoldDuration) error {
	if patron.isRegular() && patron.numberOfHolds() >= maxNumberOfHolds {
		return ErrMaxHoldsReached
	}
	return nil
}

var ErrOnlyResearcherCanPlaceOpenEndedHold = commonErrors.NewIncorrectInputError(
	"only-researcher-can-place-open-ended-hold",
	"only researcher can place open ended hold",
)

func onlyResearcherPatronsCanPlaceOpenEndedHoldsPolicy(_ book.Information, patron *Patron, duration HoldDuration) error {
	if patron.isRegular() && duration.IsOpenEnded() {
		return ErrOnlyResearcherCanPlaceOpenEndedHold
	}
	return nil
}
