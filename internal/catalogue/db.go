package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"reflect"

	"github.com/ThreeDotsLabs/watermill"
	watermillSQL "github.com/ThreeDotsLabs/watermill-sql/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/chiennguyen196/go-library/internal/catalogue/models"
	"github.com/chiennguyen196/go-library/internal/common/database"
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

func (d DB) AddABookInstance(ctx context.Context, instance BookInstance, event BookInstanceAdded) error {
	i := models.BookInstance{
		BookID:          instance.bookID,
		BookIsbn:        instance.bookIsbn,
		LibraryBranchID: instance.libraryBranchID,
		BookType:        instance.bookType,
	}
	return database.WithTx(ctx, d.db, func(tx *sql.Tx) error {
		if err := i.Insert(ctx, d.db, boil.Infer()); err != nil {
			return errors.Wrap(err, "insert book instance")
		}

		if err := publishEvents(tx, event); err != nil {
			return errors.Wrap(err, "publish event")
		}
		return nil
	})
}

func publishEvents(tx *sql.Tx, events ...interface{}) error {
	publisher, err := watermillSQL.NewPublisher(
		tx,
		watermillSQL.PublisherConfig{
			SchemaAdapter: watermillSQL.DefaultPostgreSQLSchema{},
		},
		logger,
	)
	if err != nil {
		return errors.Wrap(err, "create publisher")
	}

	messages := make(message.Messages, 0, len(events))
	for _, e := range events {
		payload, err := json.Marshal(e)
		if err != nil {
			return errors.Wrap(err, "marshal event")
		}
		msg := message.NewMessage(watermill.NewUUID(), payload)
		msg.Metadata.Set("eventType", reflect.TypeOf(e).Name())
		messages = append(messages, msg)
	}

	if err := publisher.Publish(mysqlTable, messages...); err != nil {
		return errors.Wrap(err, "publish")
	}
	return nil
}
