package command

import (
	"context"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type CancelHoldHandler struct {
	patronRepository domain.PatronRepository
}

func (h CancelHoldHandler) Handle(ctx context.Context, cmd CancelHoldCommand) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	if err := h.patronRepository.Update(ctx, cmd.PatronID, func(ctx context.Context, patron *domain.Patron) error {
		return patron.CancelHold(cmd.BookID)
	}); err != nil {
		return errors.Wrap(err, "update patron")
	}
	return nil
}

type CancelHoldCommand struct {
	PatronID domain.PatronID
	BookID   domain.BookID
}

func (c CancelHoldCommand) validate() error {
	if c.PatronID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing-patron-id")
	}
	if c.BookID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing-book-id")
	}
	return nil
}
