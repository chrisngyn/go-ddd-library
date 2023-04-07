package app

import (
	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

// Application here acts like entrypoint of our service. It contains business validation logic.
// And acts like a thin layer coordinate request to appropriate domain logic.
// Each handler of it is separated to command or query handler. It flows CQRS pattern.
type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	PlaceOnHold         command.PlaceOnHoldHandler
	CancelHold          command.CancelHoldHandler
	CheckOut            command.CheckoutHandler
	ReturnBook          command.ReturnBookHandler
	MarkOverdueCheckout command.MarkOverdueCheckoutHandler
	AddNewBook          command.AddNewBookHandler
}

type Queries struct {
	PatronProfile    query.PatronProfileHandler
	ExpiredHolds     query.ExpiredHoldsHandler
	OverdueCheckouts query.OverdueCheckoutsHandler
}
