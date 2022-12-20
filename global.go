package stdlog

import (
	"fmt"
	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
	"io"
)

func SetGlobalPlatformName(name string) {
	log = log.With(constants.PLATFORM_NAME_KEY, name)
}

func SetGlobalLogLevel(level Level) {
	logLevel = level
	levelVar.Set(logLevel.Level())
}

func SetGlobalLogWriter(w io.Writer) {
	*logWriter = w
	fmt.Printf("inside set: %T\n", *logWriter)
}
