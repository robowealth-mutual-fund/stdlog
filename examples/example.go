package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	log.SetGlobalLogLevel(log.DEBUG_LEVEL)
	log.SetGlobalPlatformName("oa v2")
	log.SetEnabledReportCaller()

	optManage := &log.OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &log.JSONFieldFormatter{},
	}

	// New logger
	l := log.NewLogger(os.Stdout, optManage, log.ERROR_LEVEL)
	l.Debug("debug level")
	l.InfoWithAttrs("new logger", log.Attrs{"test_with_field": "This is a value"})

	// Global log instance
	log.Debug("debug message")
	log.DebugWithAttrs("debug message with attr(s)", log.Attrs{"debug_key_1": "value debug"})
}
