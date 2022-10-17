package stdlog

import (
	"fmt"
	"github.com/rs/zerolog"
)

type context struct {
	zerologLogger *zerolog.Logger
}

func (c *context) Str(key, val string) Context {
	zLogger := c.zerologLogger.With().Str(key, val).Logger()
	c.zerologLogger = &zLogger
	fmt.Printf("%p\n", c.zerologLogger)
	return c
}
