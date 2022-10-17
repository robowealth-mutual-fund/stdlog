package stdlog

import (
	"github.com/robowealth-mutual-fund/stdlog/constant/errorCode"
	"github.com/rs/zerolog"
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

func (e *event) Err(errCode errorCode.Interface) Event {
	e.zerologEvent.Err(errCode)
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
