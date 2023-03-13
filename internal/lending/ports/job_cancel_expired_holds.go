package ports

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func (j Job) CancelExpiredHolds(at time.Time) {
	ctx := context.Background()
	expiredHolds, err := j.app.Queries.ExpiredHolds.Handle(ctx, query.ExpiredHoldsQuery{At: at})
	if err != nil {
		log.Fatal().Err(err).Msg("List expired holds fail")
	}
	for _, hold := range expiredHolds {
		if err := j.app.Commands.CancelHold.Handle(ctx, command.CancelHoldCommand{
			PatronID: hold.PatronID,
			BookID:   hold.BookID,
		}); err != nil {
			log.Fatal().
				Err(err).
				Str("patronID", string(hold.PatronID)).
				Str("bookID", string(hold.BookID)).
				Msg("Cancel hold fail")
		}
	}
}
