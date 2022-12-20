package stdlog

import (
	"io"

	"golang.org/x/exp/slog"
)

//go:generate mockery --dir=./ --name=Logger --filename=logger.go --output=internal/mocks --outpkg=mocks
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string, err error)
	DebugWithAttrs(msg string, fields Attrs)
	InfoWithAttrs(msg string, fields Attrs)
	WarnWithAttrs(msg string, fields Attrs)
	ErrorWithAttrs(msg string, fields Attrs)
}

type Log struct {
	slogLogger slog.Logger
}

func NewLogger(w io.Writer, optManager *OptionManager, level Level) Logger {
	var attrs []slog.Attr

	if !optManager.DisabledPlatformNameKey {
		attrs = append(attrs, slog.Any(
			optManager.JSONFieldFormatter.PlatformNameKeyWithDefault(),
			optManager.PlatformName,
		))
	}

	return &Log{
		slogLogger: slog.New(slog.HandlerOptions{
			Level:       slog.Level(level),
			ReplaceAttr: loadDefaultReplaceAttr(),
		}.
			NewJSONHandler(w).
			WithAttrs(attrs)),
	}
}

func (l Log) Debug(msg string) {
	l.slogLogger.Debug(msg)
}

func (l Log) Info(msg string) {
	l.slogLogger.Info(msg)
}

func (l Log) Warn(msg string) {
	l.slogLogger.Warn(msg)
}

func (l Log) Error(msg string, err error) {
	l.slogLogger.Error(msg, err)
}

func (l Log) DebugWithAttrs(msg string, fields Attrs) {
	l.slogLogger.LogAttrs(slog.Level(DEBUG_LEVEL), msg, fields.convert()...)
}

func (l Log) InfoWithAttrs(msg string, fields Attrs) {
	l.slogLogger.LogAttrs(slog.Level(INFO_LEVEL), msg, fields.convert()...)
}

func (l Log) WarnWithAttrs(msg string, fields Attrs) {
	l.slogLogger.LogAttrs(slog.Level(WARN_LEVEL), msg, fields.convert()...)
}

func (l Log) ErrorWithAttrs(msg string, fields Attrs) {
	l.slogLogger.LogAttrs(slog.Level(ERROR_LEVEL), msg, fields.convert()...)
}
