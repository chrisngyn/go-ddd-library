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

type CheckoutHandler struct {
	patronRepository domain.PatronRepository
}

func NewCheckoutHandler(patronRepo domain.PatronRepository) CheckoutHandler {
	if patronRepo == nil {
		panic("missing patronRepo")
	}
	return CheckoutHandler{patronRepository: patronRepo}
}

func (h CheckoutHandler) Handle(ctx context.Context, cmd CheckoutCommand) (err error) {
	defer func(st time.Time) {
		monitoring.MonitorCommand(ctx, "Checkout", cmd, err, st)
	}(time.Now())

	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}
	if err := h.patronRepository.UpdateWithBook(ctx, cmd.PatronID, cmd.BookID, func(ctx context.Context, patron *patron.Patron, book *book.Book) error {
		if err := patron.Checkout(cmd.BookID); err != nil {
			return errors.Wrap(err, "patron checkout")
		}

		if err := book.Checkout(cmd.PatronID, cmd.RequestAt); err != nil {
			return errors.Wrap(err, "book checkout")
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "update patron and book")
	}
	return nil
}

type CheckoutCommand struct {
	RequestAt time.Time
	PatronID  uuid.UUID
	BookID    uuid.UUID
}

func (c CheckoutCommand) validate() error {
	if c.RequestAt.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-request-at", "missing request at")
	}
	if c.PatronID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if c.BookID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	return nil
}
