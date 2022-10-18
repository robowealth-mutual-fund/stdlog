package stdlog

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger is the global logger.
var (
	logger = New(os.Stdout)
)

type log struct {
	zerologLogger *zerolog.Logger
}

func New(w io.Writer, opts ...*Option) Logger {
	logOpt := &logOptions{}
	for _, opt := range opts {
		opt.Apply(logOpt)
	}

	zerolog.SetGlobalLevel(setLogLevel(logOpt.level))
	zerolog.TimeFieldFormat = setTimeFieldFormat(logOpt.timeFieldFormat)

	// New log internal context
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

func setLogLevel(lv Level) zerolog.Level {
	var logLv zerolog.Level

	switch lv {
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
	default:
		logLv = zerolog.NoLevel
	}

	return logLv
}

func setTimeFieldFormat(format TimeFormat) string {
	var timeFormat string

	switch format {
	case TimeFormatUnix:
		timeFormat = zerolog.TimeFormatUnix
	case TimeFormatUnixMs:
		timeFormat = zerolog.TimeFormatUnixMs
	case TimeFormatUnixMicro:
		timeFormat = zerolog.TimeFormatUnixMicro
	case TimeFormatUnixNano:
		timeFormat = zerolog.TimeFormatUnixNano
	default:
		timeFormat = time.RFC3339
	}

	return timeFormat
}
