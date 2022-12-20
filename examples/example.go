package main

import (
	log "github.com/robowealth-mutual-fund/stdlog"
	"os"
)

func main() {
	log.SetGlobalLogLevel(log.DEBUG_LEVEL)
	log.SetGlobalPlatformName("oa v2")

	optManage := &log.OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &log.JSONFieldFormatter{},
	}

	//New logger
	l := log.NewLogger(os.Stdout, optManage, log.ERROR_LEVEL)
	l.Debug("debug level")
	l.InfoWithAttrs("new logger", log.Attrs{"test_with_field": "This is a value"})

	// Global log instance
	log.Debug("some message")
	log.DebugWithAttrs("test", log.Attrs{"debug": "test"})
}
