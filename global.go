package stdlog

import (
	"github.com/rs/zerolog"
)

func SetTimestamp() {
	logger.With().Str("service", "test service")
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
