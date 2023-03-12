package service

import (
	"github.com/chiennguyen196/go-library/internal/common/database"
	"github.com/chiennguyen196/go-library/internal/lending/adapters"
	"github.com/chiennguyen196/go-library/internal/lending/app"
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
)

func NewApplication() app.Application {
	db := database.NewSqlDB()

	patronBookRepo := adapters.NewPostgresPatronBookRepository(db)
	bookRepo := adapters.NewPostgresBookRepository(db)

	return app.Application{
		Commands: app.Commands{
			PlaceOnHold: command.NewPlaceOnHoldHandler(patronBookRepo),
			CancelHold:  command.NewCancelHoldHandler(patronBookRepo),
			CheckOut:    command.NewCheckoutHandler(patronBookRepo),
			ReturnBook:  command.NewReturnBookHandler(bookRepo),
		},
	}
}
