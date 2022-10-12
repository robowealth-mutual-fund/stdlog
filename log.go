package stdlog

import (
	std "github.com/robowealth-mutual-fund/stdlog/internal/log"
)

// Logger is the global logger.
var logger = std.New()

func Debug() std.Event {
	return logger.Debug()
}

func Info() std.Event {
	return logger.Info()
}

func Error() std.Event {
	return logger.Error()
}
