package stdlog

import (
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

func IsEnabledReportCaller() bool {
	return reportCaller
}

func Debug(msg string) {
	if reportCaller {
		log.LogAttrs(DEBUG_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Debug(msg)
}

func Info(msg string) {
	if IsEnabledReportCaller() {
		log.LogAttrs(INFO_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Info(msg)
}

func Warn(msg string) {
	if IsEnabledReportCaller() {
		log.LogAttrs(WARN_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Warn(msg)
}

func Error(msg string, err error) {
	if IsEnabledReportCaller() {
		log.LogAttrs(ERROR_LEVEL.Level(), msg, Attrs{constants.CALLER_KEY: getCaller()}.convert()...)
		return
	}

	log.Error(msg, err)
}

func DebugWithAttrs(msg string, fields Attrs) {
	if IsEnabledReportCaller() {
		fields[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(DEBUG_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	log.LogAttrs(DEBUG_LEVEL.Level(), msg, fields.convert()...)
}

func InfoWithAttrs(msg string, fields Attrs) {
	if IsEnabledReportCaller() {
		fields[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(INFO_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	log.LogAttrs(INFO_LEVEL.Level(), msg, fields.convert()...)
}

func WarnWithAttrs(msg string, fields Attrs) {
	if IsEnabledReportCaller() {
		fields[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(WARN_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	log.LogAttrs(WARN_LEVEL.Level(), msg, fields.convert()...)
}

func ErrorWithAttrs(msg string, fields Attrs) {
	if IsEnabledReportCaller() {
		fields[constants.CALLER_KEY] = getCaller()
		log.LogAttrs(ERROR_LEVEL.Level(), msg, fields.convert()...)
		return
	}

	log.LogAttrs(ERROR_LEVEL.Level(), msg, fields.convert()...)
}

func loadDefaultLogger() slog.Logger {
	levelVar.Set(logLevel.Level())
	d := &slog.HandlerOptions{
		Level:       levelVar,
		ReplaceAttr: loadDefaultReplaceAttr(),
	}

	return slog.New(d.NewJSONHandler(os.Stdout))
}

func loadDefaultReplaceAttr() func(a slog.Attr) slog.Attr {
	return func(a slog.Attr) slog.Attr {
		switch a.Key {
		case constants.DEFAULT_TIMESTAMP_KEY:
			a.Key = constants.TIMESTAMP_KEY
		case constants.DEFAULT_MESSAGE_KEY:
			a.Key = constants.MESSAGE_KEY
		case constants.DEFAULT_ERROR_KEY:
			a.Key = constants.ERROR_KEY
		}

		return a
	}
}
