package stdlog

import (
	"io"
	"os"

	"golang.org/x/exp/slog"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
)

var (
	logLevel  = INFO_LEVEL
	log       = loadDefaultLogger()
	levelVar  = &slog.LevelVar{}
	logWriter = new(io.Writer)
)

func Debug2(msg string) {
	log.Debug(msg)
}

func Debug(msg string) {
	log.Debug(msg)
}

func Info(msg string) {
	log.Info(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Error(msg string, err error) {
	log.Error(msg, err)
}

func DebugWithAttrs(msg string, fields Attrs) {
	log.LogAttrs(DEBUG_LEVEL.Level(), msg, fields.convert()...)
}

func InfoWithAttrs(msg string, fields Attrs) {
	log.LogAttrs(INFO_LEVEL.Level(), msg, fields.convert()...)
}

func WarnWithAttrs(msg string, fields Attrs) {
	log.LogAttrs(WARN_LEVEL.Level(), msg, fields.convert()...)
}

func ErrorWithAttrs(msg string, fields Attrs) {
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
		}

		return a
	}
}
