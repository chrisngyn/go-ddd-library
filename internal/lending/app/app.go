package app

import (
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	PlaceOnHold command.PlaceOnHoldHandler
	CancelHold  command.CancelHoldHandler
	CheckOut    command.CheckoutHandler
	ReturnBook  command.ReturnBookHandler
}

type Queries struct {
	PatronProfile query.PatronProfileHandler
	ExpiredHolds  query.ExpiredHoldsHandler
}
