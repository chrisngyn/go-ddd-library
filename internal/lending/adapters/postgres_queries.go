package adapters

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func getPatronByID(ctx context.Context, executor boil.ContextExecutor, patronID domain.PatronID, forUpdate bool) (domain.Patron, error) {
	mods := []qm.QueryMod{
		models.PatronWhere.ID.EQ(string(patronID)),
		qm.Load(models.PatronRels.Holds),
		qm.Load(models.PatronRels.OverdueCheckouts),
	}

	if forUpdate {
		mods = append(mods, qm.For("update"))
	}

	patron, err := models.Patrons(mods...).One(ctx, executor)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Patron{}, domain.ErrPatronNotFound
	}
	if err != nil {
		return domain.Patron{}, err
	}

	return toPatronDomain(patron)
}

func getBookByID(ctx context.Context, executor boil.ContextExecutor, bookID domain.BookID, forUpdate bool) (domain.Book, error) {
	mods := []qm.QueryMod{
		models.BookWhere.ID.EQ(string(bookID)),
	}
	if forUpdate {
		mods = append(mods, qm.For("update"))
	}

	book, err := models.Books(mods...).One(ctx, executor)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Book{}, domain.ErrBookNotFound
	}
	if err != nil {
		return domain.Book{}, err
	}

	return toDomainBook(book)
}

func updatePatron(ctx context.Context, executor boil.ContextExecutor, patron domain.Patron) error {
	if err := updateHolds(ctx, executor, patron); err != nil {
		return errors.Wrap(err, "update holds")
	}
	if err := updateOverdueCheckouts(ctx, executor, patron); err != nil {
		return errors.Wrap(err, "update overdue checkouts")
	}
	return nil
}

func updateHolds(ctx context.Context, executor boil.ContextExecutor, patron domain.Patron) error {
	_, err := models.Holds(
		models.HoldWhere.PatronID.EQ(string(patron.ID())),
	).DeleteAll(ctx, executor, false)
	if err != nil {
		return errors.Wrap(err, "delete holds")
	}

	for _, h := range patron.Holds() {
		hold := models.Hold{
			PatronID:        string(patron.ID()),
			BookID:          string(h.BookID),
			LibraryBranchID: string(h.PlacedAt),
			HoldFrom:        h.HoldDuration.From(),
			HoldTill: null.Time{
				Time:  h.HoldDuration.Till(),
				Valid: !h.HoldDuration.Till().IsZero(),
			},
		}
		err := hold.Upsert(ctx, executor, true, []string{
			models.HoldColumns.PatronID,
			models.HoldColumns.BookID,
		}, boil.Whitelist(models.HoldColumns.DeletedAt), boil.Infer())
		if err != nil {
			return errors.Wrap(err, "upsert hold")
		}
	}
	return nil
}

func updateOverdueCheckouts(ctx context.Context, executor boil.ContextExecutor, patron domain.Patron) error {
	_, err := models.OverdueCheckouts(
		models.OverdueCheckoutWhere.PatronID.EQ(string(patron.ID())),
	).DeleteAll(ctx, executor, false)
	if err != nil {
		return errors.Wrap(err, "delete overdue checkouts")
	}

	for libraryBranch, books := range patron.OverdueCheckouts() {
		for _, b := range books {
			overdueCheckout := models.OverdueCheckout{
				PatronID:        string(patron.ID()),
				BookID:          string(b),
				LibraryBranchID: string(libraryBranch),
				DeletedAt:       null.Time{},
			}
			err := overdueCheckout.Upsert(ctx, executor, true, []string{
				models.OverdueCheckoutColumns.PatronID,
				models.OverdueCheckoutColumns.BookID,
			}, boil.Whitelist(models.OverdueCheckoutColumns.DeletedAt), boil.Infer())
			if err != nil {
				return errors.Wrap(err, "upsert overdue checkout")
			}
		}
	}
	return nil
}

func updateBook(ctx context.Context, executor boil.ContextExecutor, book domain.Book) error {
	_, err := models.Books(models.BookWhere.ID.EQ(string(book.BookInfo().BookID))).
		UpdateAll(ctx, executor, models.M{
			models.BookColumns.BookStatus: toDBBookStatus(book.Status()),
			models.BookColumns.PatronID:   null.NewString(string(book.ByPatronID()), book.ByPatronID() != ""),
			models.BookColumns.HoldTill: null.NewTime(
				book.BookHoldInfo().Till,
				!book.BookHoldInfo().Till.IsZero(),
			),
			models.BookColumns.CheckedOutAt: null.NewTime(
				book.BookCheckedOutInfo().At,
				!book.BookCheckedOutInfo().At.IsZero(),
			),
		})
	return err
}
