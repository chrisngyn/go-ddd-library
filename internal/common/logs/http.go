package logs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func NewHTTPStructuredLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&HTTPStructuredLogger{logger})
}

// HTTPStructuredLogger based on example from chi: https://github.com/go-chi/chi/blob/master/_examples/logging/main.go
type HTTPStructuredLogger struct {
	Logger zerolog.Logger
}

func (l *HTTPStructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	logger := l.Logger.With().
		Str("req_id", middleware.GetReqID(r.Context())).
		Logger()

	// update logger into context
	*r = *(r.WithContext(logger.WithContext(r.Context())))

	entry := &StructuredLoggerEntry{Logger: logger}

	entry.Logger.Info().
		Str("http_method", r.Method).
		Str("uri", r.RequestURI).
		Str("remote_addr", r.RemoteAddr).
		Msg("Request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger zerolog.Logger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, _ http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger.Info().
		Int("resp_status", status).
		Int("resp_bytes_length", bytes).
		Dur("resp_elapsed", elapsed).
		Msg("Request completed")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, _ []byte) {
	l.Logger = l.Logger.With().
		Stack().
		Str("panic", fmt.Sprintf("%+v", v)).
		Logger()
}

func GetLogEntry(r *http.Request) zerolog.Logger {
	entry := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	return entry.Logger
}
