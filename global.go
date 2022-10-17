package stdlog

import (
	"github.com/rs/zerolog"
)

func SetGlobalTimestamp() {
	logger.With().Timestamp()
}

func SetGlobalPlatformName(name string) {
	logger.With().Str("platform_name", name)
}

func SetTimestampFieldName(key string) {
	zerolog.TimestampFieldName = key
}

func SetLevelFieldName(key string) {
	zerolog.LevelFieldName = key
}

func SetMessageFieldName(key string) {
	zerolog.MessageFieldName = key
}
