package logs

import (
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.New(os.Stderr)

	if isLocalEnv, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocalEnv {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	logger = logger.
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	// set global logger
	log.Logger = logger
	zerolog.DefaultContextLogger = &logger
}
