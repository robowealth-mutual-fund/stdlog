package log

import (
	"github.com/rs/zerolog"
)

type Event interface {
	Msg(msg string)
	Msgf(format string, v ...interface{})
	Interface(key string, v ...interface{})
}

type event struct {
	zerologEvent *zerolog.Event
}

func (e *event) Msg(msg string) {
	e.zerologEvent.Msg(msg)
}

func (e *event) Msgf(format string, v ...interface{}) {
	e.zerologEvent.Msgf(format, v)
}

func (e *event) Interface(key string, v ...interface{}) {
	e.zerologEvent.Interface(key, v)
}
