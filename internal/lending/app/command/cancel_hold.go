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

func NewCancelHoldHandler(patronRepo domain.PatronRepository) CancelHoldHandler {
	if patronRepo == nil {
		panic("missing patronRepo")
	}
	return CancelHoldHandler{patronRepository: patronRepo}
}

func (h CancelHoldHandler) Handle(ctx context.Context, cmd CancelHoldCommand) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	if err := h.patronRepository.UpdateWithBook(ctx, cmd.PatronID, cmd.BookID, func(ctx context.Context, patron *domain.Patron, book *domain.Book) error {
		if err := patron.CancelHold(cmd.BookID); err != nil {
			return errors.Wrap(err, "patron cancel hold")
		}
		if err := book.CancelHold(); err != nil {
			return errors.Wrap(err, "book cancel hold")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "update patron and book")
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
