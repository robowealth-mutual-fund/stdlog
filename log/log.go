package log

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info() Event
	Debug() Event
	Error() Event
}

type log struct {
	zerologLogger *zerolog.Logger
}

func (l *log) Info() Event {

	return &event{zerologEvent: l.zerologLogger.Info()}
}

func (l *log) Debug() Event {
	return &event{zerologEvent: l.zerologLogger.Debug()}
}

func (l *log) Error() Event {
	return &event{zerologEvent: l.zerologLogger.Error()}
}

type logOptions struct {
	level Level
}

type Option struct {
	fn func(*logOptions)
}

func (opt *Option) Apply(do *logOptions) {
	opt.fn(do)
}

func newLogOption(fn func(*logOptions)) Option {
	return Option{
		fn: fn,
	}
}

func WithGlobalLevel(level Level) Option {
	return newLogOption(func(o *logOptions) {
		o.level = level
	})
}

func New(w io.Writer, opts ...*Option) Logger {
	logOpt := &logOptions{}
	for _, opt := range opts {
		opt.Apply(logOpt)
	}

	var logLv zerolog.Level
	switch logOpt.level {
	case TraceLevel:
		logLv = zerolog.TraceLevel
	case DebugLevel:
		logLv = zerolog.DebugLevel
	case InfoLevel:
		logLv = zerolog.InfoLevel
	case WarnLevel:
		logLv = zerolog.WarnLevel
	case ErrorLevel:
		logLv = zerolog.ErrorLevel
	case FatalLevel:
		logLv = zerolog.FatalLevel
	case PanicLevel:
		logLv = zerolog.PanicLevel
	case Disabled:
		logLv = zerolog.Disabled
	case NoLevel:
		logLv = zerolog.NoLevel
	}

	zerolog.SetGlobalLevel(logLv)

	l := zerolog.New(w)

	return &log{
		zerologLogger: &l,
	}
}
