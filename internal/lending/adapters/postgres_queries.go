package adapters

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
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
		},
			boil.Whitelist(
				models.HoldColumns.DeletedAt,
				models.HoldColumns.HoldFrom,
				models.HoldColumns.HoldTill,
				models.HoldColumns.LibraryBranchID,
			),
			boil.Infer(),
		)
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

/*
helper functions convert db entities to domain models
*/

func toPatronDomain(patron *models.Patron) (domain.Patron, error) {
	patronType, err := toDomainPatronType(patron.PatronType)
	if err != nil {
		return domain.Patron{}, errors.Wrap(err, "to patron type")
	}

	domainHolds, err := toDomainHolds(patron.R.Holds)
	if err != nil {
		return domain.Patron{}, errors.Wrap(err, "to domain holds")
	}

	domainPatron, err := domain.NewPatron(
		domain.PatronID(patron.ID),
		patronType,
		domainHolds,
		toDomainOverdueCheckouts(patron.R.OverdueCheckouts),
	)
	if err != nil {
		return domain.Patron{}, commonErrors.NewUnknownErr("invalid-patron", err.Error())
	}
	return domainPatron, nil
}

func toDomainPatronType(patronType models.PatronType) (domain.PatronType, error) {
	switch patronType {
	case models.PatronTypeRegular:
		return domain.PatronTypeRegular, nil
	case models.PatronTypeResearcher:
		return domain.PatronTypeResearcher, nil
	default:
		return domain.PatronType{},
			commonErrors.NewUnknownErr("invalid-patron-type", fmt.Sprintf("unknown patron type %s", patronType))
	}
}

func toDomainHolds(holds models.HoldSlice) ([]domain.Hold, error) {
	domainHolds := make([]domain.Hold, 0, len(holds))
	for _, h := range holds {
		holdDuration, err := domain.NewHoldDurationFromTill(h.HoldFrom, h.HoldTill.Time)
		if err != nil {
			return nil, commonErrors.NewUnknownErr("invalid-hold-duration", err.Error())
		}
		domainHold, err := domain.NewHold(domain.BookID(h.BookID), domain.LibraryBranchID(h.LibraryBranchID), holdDuration)
		if err != nil {
			return nil, commonErrors.NewUnknownErr("invalid-hold", err.Error())
		}
		domainHolds = append(domainHolds, domainHold)
	}
	return domainHolds, nil
}

func toDomainOverdueCheckouts(overCheckouts models.OverdueCheckoutSlice) map[domain.LibraryBranchID][]domain.BookID {
	result := make(map[domain.LibraryBranchID][]domain.BookID, len(overCheckouts))
	for _, oc := range overCheckouts {
		libraryBranchID := domain.LibraryBranchID(oc.LibraryBranchID)
		bookID := domain.BookID(oc.BookID)
		result[libraryBranchID] = append(result[libraryBranchID], bookID)
	}
	return result
}

func toDomainBook(book *models.Book) (domainBook domain.Book, err error) {
	bookType, err := toDomainBookType(book.BookType)
	if err != nil {
		return domainBook, errors.Wrap(err, "to domain book type")
	}

	bookInfo, err := domain.NewBookInformation(
		domain.BookID(book.ID),
		bookType,
		domain.LibraryBranchID(book.LibraryBranchID),
	)
	if err != nil {
		return domainBook, commonErrors.NewUnknownErr("invalid-book-information", "invalid book information")
	}

	switch book.BookStatus {
	case models.BookStatusAvailable:
		return domain.NewAvailableBook(bookInfo)
	case models.BookStatusOnHold:
		return domain.NewBookOnHold(bookInfo, domain.HoldInformation{
			ByPatron: domain.PatronID(book.PatronID.String),
			Till:     book.HoldTill.Time,
		})
	case models.BookStatusCheckedOut:
		return domain.NewCheckedOutBook(bookInfo, domain.CheckedOutInformation{
			ByPatron: domain.PatronID(book.PatronID.String),
			At:       book.CheckedOutAt.Time,
		})
	default:
		return domainBook, commonErrors.NewUnknownErr("invalid-book-status", "invalid book status")
	}
}

func toDomainBookType(bookType models.BookType) (domain.BookType, error) {
	switch bookType {
	case models.BookTypeRestricted:
		return domain.BookTypeRestricted, nil
	case models.BookTypeCirculating:
		return domain.BookTypeCirculating, nil
	default:
		return domain.BookType{}, commonErrors.NewUnknownErr(
			"invalid-book-type",
			fmt.Sprintf("invalid book type %s", bookType),
		)
	}
}

func toDBBookStatus(status domain.BookStatus) models.BookStatus {
	switch status {
	case domain.BookStatusAvailable:
		return models.BookStatusAvailable
	case domain.BookStatusOnHold:
		return models.BookStatusOnHold
	case domain.BookStatusCheckedOut:
		return models.BookStatusCheckedOut
	default:
		return ""
	}
}

func toDBBookType(bookType domain.BookType) models.BookType {
	switch bookType {
	case domain.BookTypeCirculating:
		return models.BookTypeCirculating
	case domain.BookTypeRestricted:
		return models.BookTypeRestricted
	default:
		return ""
	}
}
