package stdlog

import (
	"github.com/rs/zerolog"
	"time"
)

func SetGlobalTimestamp() {
	logger.With().Timestamp()
}

func SetGlobalLogLevel(level Level) {
	var logLv zerolog.Level

	switch level {
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

	zerolog.SetGlobalLevel(logLv)
}

func SetGlobalPlatformName(name string) {
	logger.With().Str("platform_name", name)
}

func SetGlobalTimestampFieldName(key string) {
	zerolog.TimestampFieldName = key
}

func SetGlobalTimeFieldFormat(format TimeFormat) {
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

	zerolog.TimeFieldFormat = timeFormat
}

func SetGlobalLevelFieldName(key string) {
	zerolog.LevelFieldName = key
}

func SetGlobalMessageFieldName(key string) {
	zerolog.MessageFieldName = key
}
