package stdlog

import (
	"github.com/rs/zerolog"

	errorCode "github.com/robowealth-mutual-fund/stdlog/constant/errors"
)

//go:generate mockery --dir=./ --name=Logger --filename=logger.go --output=internal/mocks --outpkg=mocks
//go:generate mockery --dir=./ --name=Context --filename=context.go --output=internal/mocks --outpkg=mocks
//go:generate mockery --dir=./ --name=Event --filename=event.go --output=internal/mocks --outpkg=mocks

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
	Errs(key string, errs []errorCode.Interface) Event
	Str(key, val string) Event
	Strs(key string, vals []string) Event
}

type Context interface {
	Str(key, val string) Context
	Timestamp() Context
}
