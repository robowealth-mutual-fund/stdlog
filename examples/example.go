package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	// New log instance
	fluentLog := log.New(os.Stderr,
		log.WithPlatformName("Odini"),
		log.WithGlobalLevel(log.WarnLevel),
		log.WithTimestamp(true),
	)
	fluentLog.Warn().Str("some_key", "Hello World").Msg("This is standard log lib!")

	// Global log instance
	log.SetGlobalTimestamp()
	log.SetGlobalPlatformName("Finvest")
	log.SetTimestampFieldName("timestamp")
	log.Warn().Msg("Hello test test")
}
