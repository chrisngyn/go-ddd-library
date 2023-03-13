package domain

import (
	"context"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

var (
	ErrPatronNotFound = commonErrors.NewIncorrectInputError("patron-not-found", "patron not found")
	ErrBookNotFound   = commonErrors.NewIncorrectInputError("book-not-found", "book not found")
)

type PatronRepository interface {
	Update(
		ctx context.Context,
		patronID PatronID,
		updateFn func(ctx context.Context, patron *Patron) error,
	) error
	UpdateWithBook(
		ctx context.Context,
		patronID PatronID,
		bookID BookID,
		updateFn func(ctx context.Context, patron *Patron, book *Book) error,
	) error
}

type BookRepository interface {
	Update(
		ctx context.Context,
		bookID BookID,
		updateFn func(ctx context.Context, book *Book) error,
	) error

	UpdateWithPatron(
		ctx context.Context,
		bookID BookID,
		updateFn func(ctx context.Context, book *Book, patron *Patron) error,
	) error
}
