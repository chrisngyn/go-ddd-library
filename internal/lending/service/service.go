// Package service is where we initialize the application and its dependencies.
// At here, we will wire up all the dependencies and create the application instance.
// It also contains the component tests, and it's separated with go build tag component.
package service

import (
	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func NewApplication() (app.Application, func()) {
	db := database.NewSqlDB()

	patronRepo := adapters.NewPostgresPatronRepository(db)
	bookRepo := adapters.NewPostgresBookRepository(db)

	anApp := app.Application{
		Commands: app.Commands{
			PlaceOnHold:         command.NewPlaceOnHoldHandler(patronRepo),
			CancelHold:          command.NewCancelHoldHandler(patronRepo),
			CheckOut:            command.NewCheckoutHandler(patronRepo),
			ReturnBook:          command.NewReturnBookHandler(bookRepo),
			MarkOverdueCheckout: command.NewMarkOverdueCheckoutHandler(patronRepo),
			AddNewBook:          command.NewAddNewBookHandler(bookRepo),
		},
		Queries: app.Queries{
			PatronProfile:    query.NewPatronProfileHandler(patronRepo),
			ExpiredHolds:     query.NewExpiredHoldsHandler(bookRepo),
			OverdueCheckouts: query.NewOverdueCheckoutsHandler(bookRepo),
		},
	}

	return anApp, func() {
		_ = db.Close()
	}
}
