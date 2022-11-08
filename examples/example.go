package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	// Set global platform name value (apply only global log instance)
	// log.NewGlobalSetting().
	//	WithLevel(log.ErrorLevel).
	//	WithPlatformName("Odini").
	//	Configure()

	jsonFieldFormatter := log.NewJSONFieldFormatter()
	optManage := log.NewOptionManager(jsonFieldFormatter).WithPlatformName("Finvest")
	lg := log.NewLogger(os.Stdout, optManage)

	//New logger
	lg.Debug("debug level")
	lg.LogAttrs(log.InfoLevel, "new logger", log.Attrs{"test with field": "This is a value"})

	// Global log instance
	log.Error("some message", nil)
}
