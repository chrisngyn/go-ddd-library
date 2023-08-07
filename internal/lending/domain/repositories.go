package domain

import (
	"context"

	"github.com/google/uuid"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

var (
	ErrPatronNotFound = commonErrors.NewIncorrectInputError("patron-not-found", "patron not found")
	ErrBookNotFound   = commonErrors.NewIncorrectInputError("book-not-found", "book not found")
)

type PatronRepository interface {
	Update(
		ctx context.Context,
		patronID uuid.UUID,
		updateFn func(ctx context.Context, patron *patron.Patron) error,
	) error
	UpdateWithBook(
		ctx context.Context,
		patronID uuid.UUID,
		bookID uuid.UUID,
		updateFn func(ctx context.Context, patron *patron.Patron, book *book.Book) error,
	) error
}

type BookRepository interface {
	CreateAvailableBook(ctx context.Context, book book.Information) error
	Update(
		ctx context.Context,
		bookID uuid.UUID,
		updateFn func(ctx context.Context, book *book.Book) error,
	) error

	UpdateWithPatron(
		ctx context.Context,
		bookID uuid.UUID,
		updateFn func(ctx context.Context, book *book.Book, patron *patron.Patron) error,
	) error
}
