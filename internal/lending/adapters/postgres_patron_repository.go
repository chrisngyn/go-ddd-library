package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

type PostgresPatronRepository struct {
	db *sql.DB
}

func NewPostgresPatronRepository(db *sql.DB) PostgresPatronRepository {
	if db == nil {
		panic("missing db")
	}
	return PostgresPatronRepository{db: db}
}

func (r PostgresPatronRepository) Get(ctx context.Context, patronID uuid.UUID) (patron.Patron, error) {
	return getPatronByID(ctx, r.db, patronID, false)
}

func (r PostgresPatronRepository) Update(ctx context.Context, patronID uuid.UUID, updateFn func(ctx context.Context, patron *patron.Patron) error) error {
	return database.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		aPatron, err := getPatronByID(ctx, tx, patronID, true)
		if err != nil {
			return errors.Wrap(err, "get patron")
		}

		if err := updateFn(ctx, &aPatron); err != nil {
			return err
		}
		if err := updatePatron(ctx, tx, aPatron); err != nil {
			return errors.Wrap(err, "update patron")
		}
		return nil
	})
}

func (r PostgresPatronRepository) UpdateWithBook(ctx context.Context, patronID uuid.UUID, bookID uuid.UUID, updateFn func(ctx context.Context, patron *patron.Patron, book *book.Book) error) error {
	return database.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		aPatron, err := getPatronByID(ctx, tx, patronID, true)
		if err != nil {
			return errors.Wrap(err, "get patron")
		}
		aBook, err := getBookByID(ctx, tx, bookID, true)
		if err != nil {
			return errors.Wrap(err, "get book")
		}

		if err := updateFn(ctx, &aPatron, &aBook); err != nil {
			return errors.Wrap(err, "update")
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

func (r PostgresPatronRepository) GetPatronProfile(ctx context.Context, patronID uuid.UUID) (p query.PatronProfile, err error) {
	aPatron, err := models.Patrons(
		models.PatronWhere.ID.EQ(patronID.String()),
		qm.Load(models.PatronRels.Holds),
		qm.Load(models.PatronRels.OverdueCheckouts),
	).One(ctx, r.db)
	if errors.Is(err, sql.ErrNoRows) {
		return p, domain.ErrPatronNotFound
	}
	if err != nil {
		return p, errors.Wrap(err, "get patron")
	}
	books, err := models.Books(
		models.BookWhere.BookStatus.EQ(models.BookStatusCheckedOut),
		models.BookWhere.PatronID.EQ(null.StringFrom(patronID.String())),
	).All(ctx, r.db)
	if err != nil {
		return p, errors.Wrap(err, "get checkout books")
	}

	return toQueryPatronProfile(aPatron, books)
}

func toQueryPatronProfile(patron *models.Patron, checkedOutBooks models.BookSlice) (p query.PatronProfile, err error) {
	queryCheckedOuts := make([]query.CheckedOut, 0, len(checkedOutBooks))
	for _, c := range checkedOutBooks {
		queryCheckedOuts = append(queryCheckedOuts, query.CheckedOut{
			BookID:          uuid.MustParse(c.ID),
			LibraryBranchID: uuid.MustParse(c.LibraryBranchID),
			At:              c.CheckedOutAt.Time,
		})
	}
	patronType, err := toDomainPatronType(patron.PatronType)
	if err != nil {
		return p, err
	}
	holds, err := toDomainHolds(patron.R.Holds)
	if err != nil {
		return p, err
	}

	queryOverdueCheckouts := make([]query.OverdueCheckout, 0, len(patron.R.OverdueCheckouts))
	for _, c := range patron.R.OverdueCheckouts {
		queryOverdueCheckouts = append(queryOverdueCheckouts, query.OverdueCheckout{
			PatronID:        uuid.MustParse(patron.ID),
			BookID:          uuid.MustParse(c.BookID),
			LibraryBranchID: uuid.MustParse(c.LibraryBranchID),
		})
	}

	return query.PatronProfile{
		PatronID:         uuid.MustParse(patron.ID),
		PatronType:       patronType,
		Holds:            holds,
		CheckedOuts:      queryCheckedOuts,
		OverdueCheckouts: queryOverdueCheckouts,
	}, nil
}
