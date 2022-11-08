package main

import (
	log "github.com/robowealth-mutual-fund/stdlog"
	"os"
)

func main() {
	// Set global platform name value (apply only global log instance)
	//log.NewGlobalSetting().
	//	WithLevel(log.ErrorLevel).
	//	WithPlatformName("Odini").
	//	Build()

	jsonFieldFormatter := log.NewJSONFieldFormatter()
	optManage := log.NewOptionManager(jsonFieldFormatter).WithPlatformName("Finvest")
	lg := log.NewLogger(os.Stdout, optManage)

	lg.Debug("sss")
	//New logger
	lg.LogAttrs("new logger", log.Fields{"test with field": "This is a value"})

	// Global log instance
	log.Error("some message", nil)
}
