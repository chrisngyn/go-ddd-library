package main

import (
	"time"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"

	"github.com/chiennguyen196/go-library/internal/catalogue/models"
)

type Book struct {
	isbn   string
	title  string
	author string
}

func NewBook(isbn, title, author string) (Book, error) {
	if isbn == "" {
		return Book{}, commonErrors.NewIncorrectInputError("missing-isbn", "missing isbn")
	}
	if title == "" {
		return Book{}, commonErrors.NewIncorrectInputError("missing-title", "missing title")
	}
	if author == "" {
		return Book{}, commonErrors.NewIncorrectInputError("missing-author", "missing author")
	}
	return Book{
		isbn:   isbn,
		title:  title,
		author: author,
	}, nil
}

type BookInstance struct {
	bookID          string
	bookIsbn        string
	libraryBranchID string
	bookType        models.BookType
}

func NewBookInstance(bookID, bookIsbn, libraryBranchID string, bookType BookType) (instance BookInstance, event BookInstanceAdded, err error) {
	if bookID == "" {
		return instance, event, commonErrors.NewIncorrectInputError("missing-book-id", "missing book id")
	}
	if bookIsbn == "" {
		return instance, event, commonErrors.NewIncorrectInputError("missing-book-isbn", "missing book isbn")
	}
	if libraryBranchID == "" {
		return instance, event, commonErrors.NewIncorrectInputError("missing-library-branch-id", "missing library branch id")
	}
	dbBookType, err := toDBBookType(bookType)
	if err != nil {
		return instance, event, commonErrors.NewIncorrectInputError("invalid-book-type", err.Error())
	}
	return BookInstance{
			bookID:          bookID,
			bookIsbn:        bookIsbn,
			libraryBranchID: libraryBranchID,
			bookType:        dbBookType,
		}, BookInstanceAdded{
			ISBN:            bookIsbn,
			BookID:          bookID,
			BookType:        dbBookType,
			LibraryBranchID: libraryBranchID,
			When:            time.Now(),
		}, nil
}

func toDBBookType(bookType BookType) (models.BookType, error) {
	switch bookType {
	case Circulating:
		return models.BookTypeCirculating, nil
	case Restricted:
		return models.BookTypeRestricted, nil
	default:
		return "", commonErrors.NewIncorrectInputError("invalid-book-type", "invalid book type")
	}
}

type BookInstanceAdded struct {
	ISBN            string          `json:"isbn"`
	BookID          string          `json:"bookID"`
	BookType        models.BookType `json:"bookType"`
	LibraryBranchID string          `json:"libraryBranchID"`
	When            time.Time       `json:"when"`
}
