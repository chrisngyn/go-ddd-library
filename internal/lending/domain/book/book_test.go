package book_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chiennguyen196/go-library/internal/lending/domain/book"
)

func TestNewBook_invalid(t *testing.T) {
	bookInfo := newExampleBookInformation(t, book.TypeCirculating)
	holdInformation := newExampleHoldInformation(t)
	checkedOutInformation := newExampleCheckoutInformation(t)

	// Available Book
	_, err := book.NewAvailableBook(book.Information{})
	assert.Error(t, err)

	// Book on hold
	_, err = book.NewBookOnHold(book.Information{}, holdInformation)
	assert.Error(t, err)

	_, err = book.NewBookOnHold(bookInfo, book.HoldInformation{})
	assert.Error(t, err)

	// Checked out book
	_, err = book.NewCheckedOutBook(book.Information{}, checkedOutInformation)
	assert.Error(t, err)

	_, err = book.NewCheckedOutBook(bookInfo, book.CheckedOutInformation{})
	assert.Error(t, err)
}

func newExampleAvailableBook(t *testing.T) book.Book {
	t.Helper()
	aBook, err := book.NewAvailableBook(newExampleBookInformation(t, book.TypeCirculating))
	if err != nil {
		t.Fatalf("Fail to create available book: %v", err)
	}
	return aBook
}

func newExampleBookOnHold(t *testing.T) book.Book {
	t.Helper()
	aBook, err := book.NewBookOnHold(newExampleBookInformation(t, book.TypeCirculating), newExampleHoldInformation(t))
	if err != nil {
		t.Fatalf("Fail to create book on hold: %v", err)
	}
	return aBook
}

func newExampleCheckedOutBook(t *testing.T) book.Book {
	t.Helper()
	aBook, err := book.NewCheckedOutBook(newExampleBookInformation(t, book.TypeCirculating), newExampleCheckoutInformation(t))
	if err != nil {
		t.Fatalf("Fail to create checked out book: %v", err)
	}
	return aBook
}

func newExampleHoldInformation(t *testing.T) book.HoldInformation {
	t.Helper()
	return book.HoldInformation{
		ByPatron: uuid.New(),
		Till:     time.Now().AddDate(0, 0, 7),
	}
}

func newExampleCheckoutInformation(t *testing.T) book.CheckedOutInformation {
	t.Helper()
	return book.CheckedOutInformation{
		ByPatron: uuid.New(),
		At:       time.Now().AddDate(0, 0, -2),
	}
}

func newExampleBookInformation(t *testing.T, bookType book.Type) book.Information {
	t.Helper()
	return book.Information{
		BookID:   uuid.New(),
		BookType: bookType,
		PlacedAt: uuid.New(),
	}
}
