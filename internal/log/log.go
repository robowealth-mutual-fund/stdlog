package log

import zeroLog "github.com/rs/zerolog/log"

type Logger interface {
	Info() Event
	Debug() Event
	Error() Event
}

type log struct{}

func NewLog() Logger {
	return &log{}
}

func (l *log) Info() Event {
	return &event{zerologEvent: zeroLog.Info()}
}

func (l *log) Debug() Event {
	return &event{zerologEvent: zeroLog.Debug()}
}

func (l *log) Error() Event {
	return &event{zerologEvent: zeroLog.Error()}
}
