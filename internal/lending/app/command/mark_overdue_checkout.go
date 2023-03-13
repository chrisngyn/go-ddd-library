package command

import (
	"context"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
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

func (h MarkOverdueCheckoutHandler) Handle(ctx context.Context, cmd MarkOverdueCheckoutCommand) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate")
	}
	return h.patronRepo.Update(ctx, cmd.PatronID, func(ctx context.Context, patron *domain.Patron) error {
		patron.MarkOverdueCheckout(cmd.BookID, cmd.LibraryBranchID)
		return nil
	})
}

type MarkOverdueCheckoutCommand struct {
	PatronID        domain.PatronID
	BookID          domain.BookID
	LibraryBranchID domain.LibraryBranchID
}

func (c MarkOverdueCheckoutCommand) validate() error {
	if c.PatronID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-patron-id", "missing patron id")
	}
	if c.BookID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if c.LibraryBranchID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-library-branch-id", "missing library branch id")
	}
	return nil
}
