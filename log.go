package stdlog

import (
	"os"

	"golang.org/x/exp/slog"
)

var (
	logLevel = InfoLevel
	log      = slog.New(slog.HandlerOptions{Level: logLevel}.
			NewJSONHandler(os.Stdout))
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

func LogAttrs(level Level, msg string, fields Attrs) {
	log.LogAttrs(slog.Level(level), msg, fields.convert()...)
}
