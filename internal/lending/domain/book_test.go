package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain"
)

func TestNewBook_invalid(t *testing.T) {
	bookInfo := newExampleBookInformation(t, domain.BookTypeCirculating)
	holdInformation := newExampleHoldInformation(t)
	checkedOutInformation := newExampleCheckoutInformation(t)

	// Available Book
	_, err := domain.NewAvailableBook(domain.BookInformation{})
	assert.Error(t, err)

	// Book on hold
	_, err = domain.NewBookOnHold(domain.BookInformation{}, holdInformation)
	assert.Error(t, err)

	_, err = domain.NewBookOnHold(bookInfo, domain.HoldInformation{})
	assert.Error(t, err)

	// Checked out book
	_, err = domain.NewCheckedOutBook(domain.BookInformation{}, checkedOutInformation)
	assert.Error(t, err)

	_, err = domain.NewCheckedOutBook(bookInfo, domain.CheckedOutInformation{})
	assert.Error(t, err)
}

func newExampleAvailableBook(t *testing.T) domain.Book {
	t.Helper()
	book, err := domain.NewAvailableBook(newExampleBookInformation(t, domain.BookTypeCirculating))
	if err != nil {
		t.Fatalf("Fail to create available book: %v", err)
	}
	return book
}

func newExampleBookOnHold(t *testing.T) domain.Book {
	t.Helper()
	book, err := domain.NewBookOnHold(newExampleBookInformation(t, domain.BookTypeCirculating), newExampleHoldInformation(t))
	if err != nil {
		t.Fatalf("Fail to create book on hold: %v", err)
	}
	return book
}

func newExampleCheckedOutBook(t *testing.T) domain.Book {
	t.Helper()
	book, err := domain.NewCheckedOutBook(newExampleBookInformation(t, domain.BookTypeCirculating), newExampleCheckoutInformation(t))
	if err != nil {
		t.Fatalf("Fail to create checked out book: %v", err)
	}
	return book
}

func newExampleHoldInformation(t *testing.T) domain.HoldInformation {
	t.Helper()
	return domain.HoldInformation{
		ByPatron: "patron1",
		Till:     time.Now().AddDate(0, 0, 7),
	}
}

func newExampleCheckoutInformation(t *testing.T) domain.CheckedOutInformation {
	t.Helper()
	return domain.CheckedOutInformation{
		ByPatron: "patron1",
		At:       time.Now().AddDate(0, 0, -2),
	}
}
