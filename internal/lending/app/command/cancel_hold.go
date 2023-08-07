package command

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/common/monitoring"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
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

func (h CancelHoldHandler) Handle(ctx context.Context, cmd CancelHoldCommand) (err error) {
	defer func(st time.Time) {
		monitoring.MonitorCommand(ctx, "CancelHold", cmd, err, st)
	}(time.Now())

	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	if err := h.patronRepository.UpdateWithBook(ctx, cmd.PatronID, cmd.BookID, func(ctx context.Context, patron *patron.Patron, book *book.Book) error {
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
	PatronID uuid.UUID
	BookID   uuid.UUID
}

func (c CancelHoldCommand) validate() error {
	if c.PatronID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing-patron-id")
	}
	if c.BookID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing-book-id")
	}
	return nil
}
