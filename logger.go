package stdlog

import (
	"context"
	"io"

	"golang.org/x/exp/slog"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
)

//go:generate mockery --dir=./ --name=Logger --filename=logger.go --output=internal/mocks --outpkg=mocks
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	DebugWithAttrs(msg string, fields Attrs)
	InfoWithAttrs(msg string, fields Attrs)
	WarnWithAttrs(msg string, fields Attrs)
	ErrorWithAttrs(msg string, err error, attrs Attrs)
}

type Log struct {
	slogLogger *slog.Logger
}

func NewLogger(w io.Writer, level Level, appName string) Logger {
	return &Log{
		slogLogger: slog.New(slog.HandlerOptions{
			Level:       level.Level(),
			ReplaceAttr: loadDefaultReplaceAttr(),
		}.NewJSONHandler(w).
			WithAttrs([]slog.Attr{slog.Any(constants.APPLICATION_NAME_KEY, appName)})),
	}
}

func (l Log) Debug(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(context.Background(), DEBUG_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Debug(msg)
}

func (l Log) Info(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(context.Background(), INFO_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Info(msg)
}

func (l Log) Warn(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(context.Background(), WARN_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Warn(msg)
}

func (l Log) Error(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(context.Background(), ERROR_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Error(msg)
}

func (l Log) DebugWithAttrs(msg string, attrs Attrs) {
	if reportCaller {
		attrs[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(context.Background(), DEBUG_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	l.slogLogger.LogAttrs(context.Background(), slog.Level(DEBUG_LEVEL), msg, attrs.convert()...)
}

func (l Log) InfoWithAttrs(msg string, attrs Attrs) {
	if reportCaller {
		attrs[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(context.Background(), INFO_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	l.slogLogger.LogAttrs(context.Background(), slog.Level(INFO_LEVEL), msg, attrs.convert()...)
}

func (l Log) WarnWithAttrs(msg string, attrs Attrs) {
	if reportCaller {
		attrs[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(context.Background(), WARN_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	l.slogLogger.LogAttrs(context.Background(), slog.Level(WARN_LEVEL), msg, attrs.convert()...)
}

func (l Log) ErrorWithAttrs(msg string, err error, attrs Attrs) {
	attrs[constants.ERROR_KEY] = err.Error()

	if reportCaller {
		attrs[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(context.Background(), ERROR_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	l.slogLogger.LogAttrs(context.Background(), slog.Level(ERROR_LEVEL), msg, attrs.convert()...)
}
