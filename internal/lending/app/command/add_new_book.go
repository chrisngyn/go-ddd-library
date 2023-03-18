package command

import (
	"context"
	"time"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/common/monitoring"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type AddNewBookHandler struct {
	bookRepo domain.BookRepository
}

func NewAddNewBookHandler(bookRepo domain.BookRepository) AddNewBookHandler {
	if bookRepo == nil {
		panic("missing bookRepo")
	}
	return AddNewBookHandler{bookRepo: bookRepo}
}

func (h AddNewBookHandler) Handle(ctx context.Context, cmd AddNewBookCommand) (err error) {
	defer func(st time.Time) {
		monitoring.MonitorCommand(ctx, "AddNewBook", cmd, err, st)
	}(time.Now())

	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	bookInfo, err := domain.NewBookInformation(cmd.BookID, cmd.BookType, cmd.LibraryBranchID)
	if err != nil {
		return errors.Wrap(err, "create book info")
	}

	if err := h.bookRepo.CreateAvailableBook(ctx, bookInfo); err != nil {
		return errors.Wrap(err, "repo create book")
	}
	return nil
}

type AddNewBookCommand struct {
	BookID          domain.BookID
	BookType        domain.BookType
	LibraryBranchID domain.LibraryBranchID
}

func (c AddNewBookCommand) validate() error {
	if c.BookID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if c.BookType.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-type", "missing book type")
	}
	if c.LibraryBranchID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-library-branch-id", "missing library branch id")
	}
	return nil
}
