package stdlog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Logger is the global logger.
var (
	logger = New(os.Stdout)
)

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

	zLogCtx := zerolog.New(w).
		With()

	if logOpt.timestampFieldName != "" {
		zerolog.TimestampFieldName = logOpt.timestampFieldName
	}

	if logOpt.platformName != "" {
		zLogCtx = zLogCtx.Str("platform_name", logOpt.platformName)
	}

	if logOpt.enableTimestamp {
		zLogCtx = zLogCtx.Timestamp()
	}

	zLog := zLogCtx.Logger()
	return &log{
		zerologLogger: &zLog,
	}
}

func Debug() Event {
	return logger.Debug()
}

func Info() Event {
	return logger.Info()
}

func Warn() Event {
	return logger.Warn()
}

func Error() Event {
	return logger.Error()
}

func Fatal() Event {
	return logger.Fatal()
}

func Panic() Event {
	return logger.Panic()
}

func (l *log) Log() *zerolog.Logger {
	return l.zerologLogger
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

func (l *log) Warn() Event {
	return &event{zerologEvent: l.zerologLogger.Warn()}
}

func (l *log) Error() Event {
	return &event{zerologEvent: l.zerologLogger.Error()}
}

func (l *log) Fatal() Event {
	return &event{zerologEvent: l.zerologLogger.Fatal()}
}

func (l *log) Panic() Event {
	return &event{zerologEvent: l.zerologLogger.Panic()}
}

func (l *log) With() Context {
	return &context{zerologLogger: l.zerologLogger}
}
