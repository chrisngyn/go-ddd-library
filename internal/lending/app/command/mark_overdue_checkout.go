package command

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/common/monitoring"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

type MarkOverdueCheckoutHandler struct {
	patronRepo domain.PatronRepository
}

func NewMarkOverdueCheckoutHandler(patronRepo domain.PatronRepository) MarkOverdueCheckoutHandler {
	if patronRepo == nil {
		panic("missing patronRepo")
	}
	return MarkOverdueCheckoutHandler{patronRepo: patronRepo}
}

func (h MarkOverdueCheckoutHandler) Handle(ctx context.Context, cmd MarkOverdueCheckoutCommand) (err error) {
	defer func(st time.Time) {
		monitoring.MonitorCommand(ctx, "MarkOverdueCheckout", cmd, err, st)
	}(time.Now())

	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate")
	}
	return h.patronRepo.Update(ctx, cmd.PatronID, func(ctx context.Context, patron *patron.Patron) error {
		patron.MarkOverdueCheckout(cmd.BookID, cmd.LibraryBranchID)
		return nil
	})
}

type MarkOverdueCheckoutCommand struct {
	PatronID        uuid.UUID
	BookID          uuid.UUID
	LibraryBranchID uuid.UUID
}

func (c MarkOverdueCheckoutCommand) validate() error {
	if c.PatronID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if c.BookID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if c.LibraryBranchID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-library-branch-id", "missing library branch id")
	}
	return nil
}
