package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	optManage := &log.OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &log.JSONFieldFormatter{},
	}
	lg := log.NewLogger(os.Stdout, optManage)

	//New logger
	lg.Debug("debug level")
	lg.LogAttrs(log.INFO_LEVEL, "new logger", log.Attrs{"test with field": "This is a value"})

	// Global log instance
	log.Error("some message", nil)
}
