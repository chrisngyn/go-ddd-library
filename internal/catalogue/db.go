package main

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/catalogue/models"
)

type DB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	if db == nil {
		panic("missing db")
	}
	return DB{db: db}
}

func (d DB) AddABook(ctx context.Context, book Book) error {
	b := models.Book{
		Isbn:   book.isbn,
		Title:  book.title,
		Author: book.author,
	}
	return b.Insert(ctx, d.db, boil.Infer())
}

func (d DB) Exists(ctx context.Context, isbn string) (bool, error) {
	return models.Books(models.BookWhere.Isbn.EQ(isbn)).Exists(ctx, d.db)
}

func (d DB) AddABookInstance(ctx context.Context, instance BookInstance) error {
	i := models.BookInstance{
		BookID:          instance.bookID,
		BookIsbn:        instance.bookIsbn,
		LibraryBranchID: instance.libraryBranchID,
		BookType:        instance.bookType,
	}
	return i.Insert(ctx, d.db, boil.Infer())
}
