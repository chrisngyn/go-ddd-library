package adapters

import (
	"fmt"

	"github.com/pkg/errors"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
	"github.com/chiennguyen196/go-library/internal/lending/adapters/models"
	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

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
