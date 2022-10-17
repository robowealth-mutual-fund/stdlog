package stdlog

import (
	"os"

	"github.com/robowealth-mutual-fund/stdlog/log"
)

// Logger is the global logger.
var logger = log.New(os.Stderr)

func Debug() log.Event {
	return logger.Debug()
}

func Info() log.Event {
	return logger.Info()
}

func Error() log.Event {
	return logger.Error()
}
