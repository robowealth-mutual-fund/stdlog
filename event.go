package stdlog

import (
	"github.com/rs/zerolog"

	"github.com/robowealth-mutual-fund/stdlog/constant/errors"
)

type event struct {
	zerologEvent *zerolog.Event
}

func (e *event) Msg(msg string) {
	e.zerologEvent.Msg(msg)
}

func (e *event) Msgf(format string, v ...interface{}) {
	e.zerologEvent.Msgf(format, v...)
}

func (e *event) Interface(key string, v ...interface{}) Event {
	e.zerologEvent.Interface(key, v)
	return e
}

func (e *event) Err(err errors.Interface) Event {
	e.zerologEvent.Err(err)
	return e
}

func (e *event) Errs(key string, errs []error) Event {
	e.zerologEvent.Errs(key, errs)
	return e
}

func (e *event) Str(key, val string) Event {
	e.zerologEvent.Str(key, val)
	return e
}

func (e *event) Strs(key string, vals []string) Event {
	e.zerologEvent.Strs(key, vals)
	return e
}
