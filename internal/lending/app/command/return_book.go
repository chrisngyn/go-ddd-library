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

type ReturnBookHandler struct {
	bookRepository domain.BookRepository
}

func NewReturnBookHandler(bookRepo domain.BookRepository) ReturnBookHandler {
	if bookRepo == nil {
		panic("missing bookRepo")
	}
	return ReturnBookHandler{bookRepository: bookRepo}
}

func (h ReturnBookHandler) Handle(ctx context.Context, cmd ReturnBookCommand) (err error) {
	defer func(st time.Time) {
		monitoring.MonitorCommand(ctx, "ReturnBook", cmd, err, st)
	}(time.Now())

	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "validate input")
	}

	if err := h.bookRepository.UpdateWithPatron(ctx, cmd.BookID, func(ctx context.Context, book *book.Book, patron *patron.Patron) error {
		if err := book.CheckIn(); err != nil {
			return errors.Wrap(err, "check in book")
		}

		patron.ReturnBook(cmd.BookID)

		return nil
	}); err != nil {
		return errors.Wrap(err, "update book with patron")
	}
	return nil
}

type ReturnBookCommand struct {
	BookID uuid.UUID
}

func (c ReturnBookCommand) validate() error {
	if c.BookID == uuid.Nil {
		return commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	return nil
}
