package monitoring

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	commonErrors "github.com/chiennguyen196/go-library/internal/common/errors"
)

func MonitorCommand(ctx context.Context, commandName string, input interface{}, err error, startTime time.Time) {
	logger := log.Ctx(ctx).With().
		Str("command", commandName).
		Dur("elapsed", time.Since(startTime)).
		Interface("input", input).
		Logger()

	if err == nil {
		logger.Info().Msg("Execute command succeed")
		return
	}

	var slugErr commonErrors.SlugError
	if errors.As(err, &slugErr) {
		if slugErr.ErrorType() == commonErrors.ErrorTypeIncorrectInput {
			log.Warn().
				Err(err).
				Str("slug", slugErr.Slug()).
				Msg("Incorrect input")
			return
		}
	}

	logger.Error().
		Err(err).
		Msg("Execute command failed")
}

func MonitorQuery(ctx context.Context, queryName string, input, output interface{}, err error, startTime time.Time) {
	logger := log.Ctx(ctx).With().
		Str("query", queryName).
		Dur("elapsed", time.Since(startTime)).
		Interface("input", input).
		Interface("output", output).
		Logger()

	if err == nil {
		logger.Info().Msg("Query succeed")
		return
	}

	var slugErr commonErrors.SlugError
	if errors.As(err, &slugErr) {
		if slugErr.ErrorType() == commonErrors.ErrorTypeIncorrectInput {
			log.Warn().
				Err(err).
				Str("slug", slugErr.Slug()).
				Msg("Incorrect input")
			return
		}
	}

	logger.Error().
		Err(err).
		Msg("Query failed")
}
