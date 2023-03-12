package adapters

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) PostgresBookRepository {
	if db == nil {
		panic("missing db")
	}
	return PostgresBookRepository{
		db: db,
	}
}

func (r PostgresBookRepository) Update(ctx context.Context, bookID domain.BookID, updateFn func(ctx context.Context, book *domain.Book) error) error {
	return WithTx(ctx, r.db, func(tx *sql.Tx) error {
		book, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book by id")
		}

		if err := updateFn(ctx, &book); err != nil {
			return err
		}

		if err := updateBook(ctx, tx, book); err != nil {
			return errors.Wrap(err, "update book")
		}

		return nil
	})
}

func (r PostgresBookRepository) UpdateWithPatron(ctx context.Context, bookID domain.BookID, updateFn func(ctx context.Context, book *domain.Book, patron *domain.Patron) error) error {
	return WithTx(ctx, r.db, func(tx *sql.Tx) error {
		book, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book by id")
		}
		patron, err := getPatronByID(ctx, tx, book.ByPatronID(), true)
		if err != nil {
			return errors.Wrap(err, "get patron by id")
		}

		if err := updateFn(ctx, &book, &patron); err != nil {
			return err
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
