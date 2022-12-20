package stdlog

import (
	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
	"io"
	"os"

	"golang.org/x/exp/slog"
)

var (
	logLevel  = INFO_LEVEL
	log       = loadDefaultLogger()
	levelVar  = &slog.LevelVar{}
	logWriter = new(io.Writer)
)

func Debug(msg string, args ...any) {
	log.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	log.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	log.Warn(msg, args...)
}

func Error(msg string, err error, args ...any) {
	log.Error(msg, err, args...)
}

//func LogAttrs(msg string, fields Attrs) {
//	log.LogAttrs(levelVar.Level(), msg, fields.convert()...)
//}

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
