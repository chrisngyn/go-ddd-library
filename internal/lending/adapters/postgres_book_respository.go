package adapters

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
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

func (r PostgresBookRepository) CreateAvailableBook(ctx context.Context, book book.Information) error {
	dbBook := models.Book{
		ID:              book.BookID.String(),
		LibraryBranchID: book.PlacedAt.String(),
		BookType:        toDBBookType(book.BookType),
		BookStatus:      models.BookStatusAvailable,
	}

	err := dbBook.Insert(ctx, r.db, boil.Infer())
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		if pgErr.Code.Name() == "unique_violation" {
			// Ignoring for idempotency
			log.Ctx(ctx).Info().Msg("Duplicate book id. Ignore!")
			return nil
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func (r PostgresBookRepository) Get(ctx context.Context, bookID uuid.UUID) (book.Book, error) {
	return getBookByID(ctx, r.db, bookID, false)
}

func (r PostgresBookRepository) Update(ctx context.Context, bookID uuid.UUID, updateFn func(ctx context.Context, book *book.Book) error) error {
	return database.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		aBook, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book by id")
		}

		if err := updateFn(ctx, &aBook); err != nil {
			return err
		}

		if err := updateBook(ctx, tx, aBook); err != nil {
			return errors.Wrap(err, "update book")
		}

		return nil
	})
}

func (r PostgresBookRepository) UpdateWithPatron(ctx context.Context, bookID uuid.UUID, updateFn func(ctx context.Context, book *book.Book, patron *patron.Patron) error) error {
	return database.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		aBook, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book by id")
		}
		aPatron, err := getPatronByID(ctx, tx, aBook.ByPatronID(), true)
		if err != nil {
			return errors.Wrap(err, "get patron by id")
		}

		if err := updateFn(ctx, &aBook, &aPatron); err != nil {
			return err
		}

		if err := updatePatron(ctx, tx, aPatron); err != nil {
			return errors.Wrap(err, "update patron")
		}
		if err := updateBook(ctx, tx, aBook); err != nil {
			return errors.Wrap(err, "update book")
		}
		return nil
	})
}

func (r PostgresBookRepository) ListExpiredHolds(ctx context.Context, at time.Time) ([]query.ExpiredHold, error) {
	books, err := models.Books(
		models.BookWhere.BookStatus.EQ(models.BookStatusOnHold),
		models.BookWhere.HoldTill.LTE(null.TimeFrom(at)),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "get expired hold books")
	}

	result := make([]query.ExpiredHold, 0, len(books))
	for _, b := range books {
		result = append(result, query.ExpiredHold{
			BookID:          uuid.MustParse(b.ID),
			LibraryBranchID: uuid.MustParse(b.LibraryBranchID),
			PatronID:        uuid.MustParse(b.PatronID.String),
			HoldTill:        b.HoldTill.Time,
		})
	}
	return result, nil
}

func (r PostgresBookRepository) ListOverdueCheckouts(ctx context.Context, at time.Time, maxCheckoutDurationDays int) ([]query.OverdueCheckout, error) {
	books, err := models.Books(
		models.BookWhere.BookStatus.EQ(models.BookStatusCheckedOut),
		models.BookWhere.CheckedOutAt.LTE(null.TimeFrom(at.AddDate(0, 0, -maxCheckoutDurationDays))),
	).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "list overdue checkout books")
	}

	result := make([]query.OverdueCheckout, 0, len(books))
	for _, b := range books {
		result = append(result, query.OverdueCheckout{
			PatronID:        uuid.MustParse(b.PatronID.String),
			BookID:          uuid.MustParse(b.ID),
			LibraryBranchID: uuid.MustParse(b.LibraryBranchID),
		})
	}
	return result, nil
}
