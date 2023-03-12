package command

import (
	"context"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type ReturnBookHandler struct {
	bookRepository domain.BookRepository
}

func NewReturnBookHandler(bookRepo domain.BookRepository) ReturnBookHandler {
	if bookRepo == nil {
		panic("missing bookRepo")
	}
	return ReturnBookHandler{bookRepository: bookRepo}
}

func (h ReturnBookHandler) Handle(ctx context.Context, cmd ReturnBookCommand) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	if err := h.bookRepository.UpdateWithPatron(ctx, cmd.BookID, func(ctx context.Context, book *domain.Book, patron *domain.Patron) error {
		if err := book.CheckIn(); err != nil {
			return errors.Wrap(err, "check in book")
		}

		if err := patron.ReturnBook(cmd.BookID); err != nil {
			return errors.Wrap(err, "patron return book")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "update book with patron")
	}
	return nil
}

type ReturnBookCommand struct {
	BookID domain.BookID
}

func (c ReturnBookCommand) validate() error {
	if c.BookID.IsZero() {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	return nil
}
