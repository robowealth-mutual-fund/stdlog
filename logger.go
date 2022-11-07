package stdlog

import (
	"io"

	"golang.org/x/exp/slog"
)

type Logger interface {
	Info(msg string, args ...any)
	Debug(msg string, args ...any)
	Error(msg string, err error, args ...any)
	LogAttrs(level Level, msg string, fields Fields)
	Warn(msg string, args ...any)
}

type Log struct {
	slogLogger slog.Logger
}

func NewLogger(w io.Writer, optManager *OptionManager) Logger {
	var attrs []slog.Attr

	if !optManager.DisabledPlatformNameKey() {
		attrs = append(attrs, slog.Any(
			optManager.jsonFieldManager.PlatformNameKeyWithDefault(),
			optManager.PlatformName(),
		))
	}

	return &Log{
		slogLogger: slog.New(slog.HandlerOptions{Level: logLevel}.NewJSONHandler(w).WithAttrs(attrs)),
	}
}

func (l Log) Info(msg string, args ...any) {
	l.slogLogger.Info(msg, args...)
}

func (l Log) Debug(msg string, args ...any) {
	l.slogLogger.Debug(msg, args...)
}

func (l Log) Error(msg string, err error, args ...any) {
	l.slogLogger.Error(msg, err, args...)
}

func (l Log) LogAttrs(level Level, msg string, fields Fields) {
	l.slogLogger.LogAttrs(slog.Level(level), msg, fields.convert()...)
}

func (l Log) Warn(msg string, args ...any) {
	l.slogLogger.Warn(msg, args...)
}
