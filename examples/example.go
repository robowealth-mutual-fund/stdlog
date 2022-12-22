package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	log.SetGlobalLogLevel(log.DEBUG_LEVEL)
	log.SetGlobalApplicationName("oa v2")
	log.SetEnabledReportCaller()

	//New logger
	l := log.NewLogger(os.Stdout, log.DEBUG_LEVEL, "exmaple")
	l.Debug("debug level")
	l.InfoWithAttrs("new logger", log.Attrs{"test_with_field": "This is a value"})

	// Global log instance
	log.Debug("debug message")
	log.InfoWithAttrs("debug message with attr(s)", log.Attrs{"debug_key_1": "value debug"})
}
