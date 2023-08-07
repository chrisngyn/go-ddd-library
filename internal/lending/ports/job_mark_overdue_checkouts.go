package ports

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/chiennguyen196/go-library/internal/lending/app/command"
	"github.com/chiennguyen196/go-library/internal/lending/app/query"
)

func (j Job) MarkOverdueCheckouts(at time.Time) {
	start := time.Now()
	ctx := context.Background()
	overdueCheckouts, err := j.app.Queries.OverdueCheckouts.Handle(ctx, query.OverdueCheckoutsQuery{At: at})
	if err != nil {
		log.Fatal().Err(err).Msg("List overdue checkouts fail")
	}

	for _, o := range overdueCheckouts {
		err := j.app.Commands.MarkOverdueCheckout.Handle(ctx, command.MarkOverdueCheckoutCommand{
			PatronID:        o.PatronID,
			BookID:          o.BookID,
			LibraryBranchID: o.LibraryBranchID,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Cannot mark overdue checkout")
		}
	}

	log.Info().Dur("elapsed", time.Since(start)).Msg("Done!")
}
