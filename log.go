package stdlog

import (
	"context"
	"os"

	"golang.org/x/exp/slog"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
)

var (
	logLevel     = INFO_LEVEL
	log          = loadDefaultLogger()
	levelVar     = &slog.LevelVar{}
	reportCaller bool
)

func Debug(msg string) {
	if ReportCaller() {
		log.LogAttrs(context.Background(), DEBUG_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Debug(msg)
}

func Info(msg string) {
	if ReportCaller() {
		log.LogAttrs(context.Background(), INFO_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Info(msg)
}

func Warn(msg string) {
	if ReportCaller() {
		log.LogAttrs(context.Background(), WARN_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Warn(msg)
}

func Error(msg string, err error) {
	if ReportCaller() {
		log.LogAttrs(context.Background(), ERROR_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Error(msg, err)
}

func DebugWithAttrs(msg string, attrs Attrs) {
	if ReportCaller() {
		attrs[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(context.Background(), DEBUG_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	log.LogAttrs(context.Background(), DEBUG_LEVEL.Level(), msg, attrs.convert()...)
}

func InfoWithAttrs(msg string, attrs Attrs) {
	if ReportCaller() {
		attrs[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(context.Background(), INFO_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	log.LogAttrs(context.Background(), INFO_LEVEL.Level(), msg, attrs.convert()...)
}

func WarnWithAttrs(msg string, attrs Attrs) {
	if ReportCaller() {
		attrs[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(context.Background(), WARN_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	log.LogAttrs(context.Background(), WARN_LEVEL.Level(), msg, attrs.convert()...)
}

func ErrorWithAttrs(msg string, err error, attrs Attrs) {
	attrs[constants.ERROR_KEY] = err.Error()

	if ReportCaller() {
		attrs[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(context.Background(), ERROR_LEVEL.Level(), msg, attrs.convert()...)
		return
	}

	log.LogAttrs(context.Background(), ERROR_LEVEL.Level(), msg, attrs.convert()...)
}

func loadDefaultLogger() *slog.Logger {
	levelVar.Set(logLevel.Level())
	d := &slog.HandlerOptions{
		Level:       levelVar,
		ReplaceAttr: loadDefaultReplaceAttr(),
	}

	return slog.New(d.NewJSONHandler(os.Stdout))
}

func loadDefaultReplaceAttr() func(groups []string, attr slog.Attr) slog.Attr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		switch attr.Key {
		case constants.DEFAULT_TIMESTAMP_KEY:
			attr.Key = constants.TIMESTAMP_KEY
		case constants.DEFAULT_MESSAGE_KEY:
			attr.Key = constants.MESSAGE_KEY
		case constants.DEFAULT_ERROR_KEY:
			attr.Key = constants.ERROR_KEY
		}

		return attr
	}
}
