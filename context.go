package stdlog

import (
	"github.com/rs/zerolog"
)

type context struct {
	zerologLogger *zerolog.Logger
}

func (c *context) Str(key, val string) Context {
	*c.zerologLogger = c.zerologLogger.With().Str(key, val).Logger()
	return c
}

func (c *context) Timestamp() Context {
	*c.zerologLogger = c.zerologLogger.With().Timestamp().Logger()
	return c
}
