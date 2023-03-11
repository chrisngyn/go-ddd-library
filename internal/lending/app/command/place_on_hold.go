package command

import (
	"context"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type PlaceOnHoldHandler struct {
	patronBookRepository domain.PatronBookRepository
}

func (h PlaceOnHoldHandler) Handle(ctx context.Context, cmd PlaceOnHoldCommand) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate")
	}

	if err := h.patronBookRepository.Update(ctx, cmd.PatronID, cmd.BookID, func(ctx context.Context, patron *domain.Patron, book *domain.Book) error {
		if err := patron.PlaceOnHold(book.BookInformation(), cmd.HoldDuration); err != nil {
			return errors.Wrap(err, "patron place on hold")
		}

		if err := book.HoldBy(patron.ID(), cmd.HoldDuration); err != nil {
			return errors.Wrap(err, "book hold")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "update")
	}
	return nil
}

type PlaceOnHoldCommand struct {
	PatronID     domain.PatronID
	BookID       domain.BookID
	HoldDuration domain.HoldDuration
}

func (c PlaceOnHoldCommand) validate() error {
	if c.PatronID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if c.BookID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if c.HoldDuration.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-hold-duration", "missing hold duration")
	}
	return nil
}