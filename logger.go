package stdlog

import (
	"io"

	"golang.org/x/exp/slog"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
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
	if reportCaller {
		l.slogLogger.LogAttrs(DEBUG_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Debug(msg)
}

func (l Log) Info(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(INFO_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Info(msg)
}

func (l Log) Warn(msg string) {
	if reportCaller {
		l.slogLogger.LogAttrs(WARN_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Warn(msg)
}

func (l Log) Error(msg string, err error) {
	if reportCaller {
		l.slogLogger.LogAttrs(ERROR_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	l.slogLogger.Error(msg, err)
}

func (l Log) DebugWithAttrs(msg string, fields Attrs) {
	if reportCaller {
		fields[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(DEBUG_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	l.slogLogger.LogAttrs(slog.Level(DEBUG_LEVEL), msg, fields.convert()...)
}

func (l Log) InfoWithAttrs(msg string, fields Attrs) {
	if reportCaller {
		fields[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(INFO_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	l.slogLogger.LogAttrs(slog.Level(INFO_LEVEL), msg, fields.convert()...)
}

func (l Log) WarnWithAttrs(msg string, fields Attrs) {
	if reportCaller {
		fields[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(WARN_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	l.slogLogger.LogAttrs(slog.Level(WARN_LEVEL), msg, fields.convert()...)
}

func (l Log) ErrorWithAttrs(msg string, fields Attrs) {
	if reportCaller {
		fields[constants.CALLER_KEY] = getCaller()
		l.slogLogger.LogAttrs(ERROR_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	l.slogLogger.LogAttrs(slog.Level(ERROR_LEVEL), msg, fields.convert()...)
}
