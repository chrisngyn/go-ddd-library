package service

import (
	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func NewApplication() app.Application {
	db := database.NewSqlDB()

	patronRepo := adapters.NewPostgresPatronRepository(db)
	bookRepo := adapters.NewPostgresBookRepository(db)

	return app.Application{
		Commands: app.Commands{
			PlaceOnHold: command.NewPlaceOnHoldHandler(patronRepo),
			CancelHold:  command.NewCancelHoldHandler(patronRepo),
			CheckOut:    command.NewCheckoutHandler(patronRepo),
			ReturnBook:  command.NewReturnBookHandler(bookRepo),
		},
		Queries: app.Queries{
			PatronProfile: query.NewPatronProfileHandler(patronRepo),
		},
	}
}
