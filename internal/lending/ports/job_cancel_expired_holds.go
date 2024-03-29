package ports

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func (j Job) CancelExpiredHolds(at time.Time) {
	start := time.Now()
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
				Str("patronID", hold.PatronID.String()).
				Str("bookID", hold.BookID.String()).
				Msg("Cancel hold fail")
		}
	}
	log.Info().Dur("elapsed", time.Since(start)).Msg("Done!")
}
