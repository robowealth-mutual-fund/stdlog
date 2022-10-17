package stdlog

import (
	"github.com/rs/zerolog"

	"github.com/robowealth-mutual-fund/stdlog/constant/errorCode"
)

type Logger interface {
	Info() Event
	Debug() Event
	Error() Event
	Warn() Event
	Fatal() Event
	Panic() Event
	Log() *zerolog.Logger
	With() Context
}

type Event interface {
	Msg(msg string)
	Msgf(format string, v ...interface{})
	Interface(key string, v ...interface{}) Event
	Err(errCode errorCode.Interface) Event
	Errs(key string, errs []error) Event
	Str(key, val string) Event
	Strs(key string, vals []string) Event
}

type Context interface {
	Str(key, val string) Context
}
