package logs

import (
	"context"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
)

// SQLLogAdapter based on https://github.com/simukti/sqldb-logger/blob/master/logadapter/zerologadapter/logger.go
type SQLLogAdapter struct {
}

func NewSQLLogAdapter() SQLLogAdapter {
	return SQLLogAdapter{}
}

func (a SQLLogAdapter) Log(ctx context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	var lvl zerolog.Level

	switch level {
	case sqldblogger.LevelError:
		lvl = zerolog.ErrorLevel
	case sqldblogger.LevelInfo:
		lvl = zerolog.InfoLevel
	case sqldblogger.LevelDebug:
		lvl = zerolog.DebugLevel
	case sqldblogger.LevelTrace:
		lvl = zerolog.TraceLevel
	default:
		lvl = zerolog.DebugLevel
	}

	normalizeQuery(data)

	zerolog.Ctx(ctx).WithLevel(lvl).Fields(data).Msg(msg)
}

var spaceReg, _ = regexp.Compile("[\\s\n\t]+")

func normalizeQuery(data map[string]interface{}) {
	rawQuery, ok := data["query"]
	if !ok {
		return
	}
	query, ok := rawQuery.(string)
	if !ok {
		return
	}
	data["query"] = strings.TrimSpace(spaceReg.ReplaceAllString(query, " "))
}
