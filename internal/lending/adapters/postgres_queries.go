package adapters

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
	"github.com/chiennguyen196/go-library/internal/lending/domain/patron"
)

func getPatronByID(ctx context.Context, executor boil.ContextExecutor, patronID uuid.UUID, forUpdate bool) (patron.Patron, error) {
	mods := []qm.QueryMod{
		models.PatronWhere.ID.EQ(patronID.String()),
		qm.Load(models.PatronRels.Holds),
		qm.Load(models.PatronRels.OverdueCheckouts),
	}

	if forUpdate {
		mods = append(mods, qm.For("update"))
	}

	aPatron, err := models.Patrons(mods...).One(ctx, executor)
	if errors.Is(err, sql.ErrNoRows) {
		return patron.Patron{}, domain.ErrPatronNotFound
	}
	if err != nil {
		return patron.Patron{}, err
	}

	return toPatronDomain(aPatron)
}

func getBookByID(ctx context.Context, executor boil.ContextExecutor, bookID uuid.UUID, forUpdate bool) (book.Book, error) {
	mods := []qm.QueryMod{
		models.BookWhere.ID.EQ(bookID.String()),
	}
	if forUpdate {
		mods = append(mods, qm.For("update"))
	}

	aBook, err := models.Books(mods...).One(ctx, executor)
	if errors.Is(err, sql.ErrNoRows) {
		return book.Book{}, domain.ErrBookNotFound
	}
	if err != nil {
		return book.Book{}, err
	}

	return toDomainBook(aBook)
}

func updatePatron(ctx context.Context, executor boil.ContextExecutor, patron patron.Patron) error {
	if err := updateHolds(ctx, executor, patron); err != nil {
		return errors.Wrap(err, "update holds")
	}
	if err := updateOverdueCheckouts(ctx, executor, patron); err != nil {
		return errors.Wrap(err, "update overdue checkouts")
	}
	return nil
}

func updateHolds(ctx context.Context, executor boil.ContextExecutor, patron patron.Patron) error {
	_, err := models.Holds(
		models.HoldWhere.PatronID.EQ(patron.ID().String()),
	).DeleteAll(ctx, executor, false)
	if err != nil {
		return errors.Wrap(err, "delete holds")
	}

	for _, h := range patron.Holds() {
		hold := models.Hold{
			PatronID:        patron.ID().String(),
			BookID:          h.BookID.String(),
			LibraryBranchID: h.PlacedAt.String(),
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

func updateOverdueCheckouts(ctx context.Context, executor boil.ContextExecutor, patron patron.Patron) error {
	_, err := models.OverdueCheckouts(
		models.OverdueCheckoutWhere.PatronID.EQ(patron.ID().String()),
	).DeleteAll(ctx, executor, false)
	if err != nil {
		return errors.Wrap(err, "delete overdue checkouts")
	}

	for libraryBranch, books := range patron.OverdueCheckouts() {
		for _, bid := range books {
			overdueCheckout := models.OverdueCheckout{
				PatronID:        patron.ID().String(),
				BookID:          bid.String(),
				LibraryBranchID: libraryBranch.String(),
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

func updateBook(ctx context.Context, executor boil.ContextExecutor, book book.Book) error {
	_, err := models.Books(models.BookWhere.ID.EQ(book.BookInfo().BookID.String())).
		UpdateAll(ctx, executor, models.M{
			models.BookColumns.BookStatus: toDBBookStatus(book.Status()),
			models.BookColumns.PatronID:   null.NewString(book.ByPatronID().String(), book.ByPatronID() != uuid.Nil),
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

func toPatronDomain(aPatron *models.Patron) (patron.Patron, error) {
	patronType, err := toDomainPatronType(aPatron.PatronType)
	if err != nil {
		return patron.Patron{}, errors.Wrap(err, "to patron type")
	}

	domainHolds, err := toDomainHolds(aPatron.R.Holds)
	if err != nil {
		return patron.Patron{}, errors.Wrap(err, "to domain holds")
	}

	domainPatron, err := patron.NewPatron(
		uuid.MustParse(aPatron.ID),
		patronType,
		domainHolds,
		toDomainOverdueCheckouts(aPatron.R.OverdueCheckouts),
	)
	if err != nil {
		return patron.Patron{}, commonErrors.NewUnknownErr("invalid-patron", err.Error())
	}
	return domainPatron, nil
}

func toDomainPatronType(patronType models.PatronType) (patron.Type, error) {
	switch patronType {
	case models.PatronTypeRegular:
		return patron.TypeRegular, nil
	case models.PatronTypeResearcher:
		return patron.TypeResearcher, nil
	default:
		return patron.Type{},
			commonErrors.NewUnknownErr("invalid-patron-type", fmt.Sprintf("unknown patron type %s", patronType))
	}
}

func toDomainHolds(holds models.HoldSlice) ([]patron.Hold, error) {
	domainHolds := make([]patron.Hold, 0, len(holds))
	for _, h := range holds {
		holdDuration, err := patron.NewHoldDurationFromTill(h.HoldFrom, h.HoldTill.Time)
		if err != nil {
			return nil, commonErrors.NewUnknownErr("invalid-hold-duration", err.Error())
		}
		domainHold, err := patron.NewHold(uuid.MustParse(h.BookID), uuid.MustParse(h.LibraryBranchID), holdDuration)
		if err != nil {
			return nil, commonErrors.NewUnknownErr("invalid-hold", err.Error())
		}
		domainHolds = append(domainHolds, domainHold)
	}
	return domainHolds, nil
}

func toDomainOverdueCheckouts(overCheckouts models.OverdueCheckoutSlice) patron.OverdueCheckouts {
	result := make(patron.OverdueCheckouts, len(overCheckouts))
	for _, oc := range overCheckouts {
		libraryBranchID := uuid.MustParse(oc.LibraryBranchID)
		bookID := uuid.MustParse(oc.BookID)
		result[libraryBranchID] = append(result[libraryBranchID], bookID)
	}
	return result
}

func toDomainBook(aBook *models.Book) (domainBook book.Book, err error) {
	bookType, err := toDomainBookType(aBook.BookType)
	if err != nil {
		return domainBook, errors.Wrap(err, "to domain book type")
	}

	bookInfo, err := book.NewBookInformation(
		uuid.MustParse(aBook.ID),
		bookType,
		uuid.MustParse(aBook.LibraryBranchID),
	)
	if err != nil {
		return domainBook, commonErrors.NewUnknownErr("invalid-book-information", "invalid book information")
	}

	switch aBook.BookStatus {
	case models.BookStatusAvailable:
		return book.NewAvailableBook(bookInfo)
	case models.BookStatusOnHold:
		return book.NewBookOnHold(bookInfo, book.HoldInformation{
			ByPatron: uuid.MustParse(aBook.PatronID.String),
			Till:     aBook.HoldTill.Time,
		})
	case models.BookStatusCheckedOut:
		return book.NewCheckedOutBook(bookInfo, book.CheckedOutInformation{
			ByPatron: uuid.MustParse(aBook.PatronID.String),
			At:       aBook.CheckedOutAt.Time,
		})
	default:
		return domainBook, commonErrors.NewUnknownErr("invalid-book-status", "invalid book status")
	}
}

func toDomainBookType(bookType models.BookType) (book.Type, error) {
	switch bookType {
	case models.BookTypeRestricted:
		return book.TypeRestricted, nil
	case models.BookTypeCirculating:
		return book.TypeCirculating, nil
	default:
		return book.Type{}, commonErrors.NewUnknownErr(
			"invalid-book-type",
			fmt.Sprintf("invalid book type %s", bookType),
		)
	}
}

func toDBBookStatus(status book.Status) models.BookStatus {
	switch status {
	case book.StatusAvailable:
		return models.BookStatusAvailable
	case book.StatusOnHold:
		return models.BookStatusOnHold
	case book.StatusCheckedOut:
		return models.BookStatusCheckedOut
	default:
		return ""
	}
}

func toDBBookType(bookType book.Type) models.BookType {
	switch bookType {
	case book.TypeCirculating:
		return models.BookTypeCirculating
	case book.TypeRestricted:
		return models.BookTypeRestricted
	default:
		return ""
	}
}
