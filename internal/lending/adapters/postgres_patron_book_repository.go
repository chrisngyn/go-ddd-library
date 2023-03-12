package adapters

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type PostgresPatronBookRepository struct {
	db *sql.DB
}

func NewPostgresPatronBookRepository(db *sql.DB) PostgresPatronBookRepository {
	if db == nil {
		panic("missing db")
	}
	return PostgresPatronBookRepository{db: db}
}

func (p PostgresPatronBookRepository) Update(ctx context.Context, patronID domain.PatronID, bookID domain.BookID, updateFn func(ctx context.Context, patron *domain.Patron, book *domain.Book) error) error {
	return WithTx(ctx, p.db, func(tx *sql.Tx) error {
		patron, err := getPatronByID(ctx, tx, patronID, true)
		if err != nil {
			return errors.Wrap(err, "get patron")
		}
		book, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book")
		}

		if err := updateFn(ctx, &patron, &book); err != nil {
			return errors.Wrap(err, "update")
		}

		if err := updatePatron(ctx, tx, patron); err != nil {
			return errors.Wrap(err, "update patron")
		}
		if err := updateBook(ctx, tx, book); err != nil {
			return errors.Wrap(err, "update book")
		}

		return nil
	})
}
