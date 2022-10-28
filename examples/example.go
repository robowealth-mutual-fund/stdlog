package main

import (
	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	// New log instance
	//fluentLog := log.New(os.Stderr,
	//	log.WithPlatformName("Odini"),
	//	log.WithTimestamp(true),
	//)
	//fluentLog.Warn().Str("some_key", "Hello World").Msg("This is standard log lib!")
	//
	//// Global log instance
	//log.SetGlobalTimestamp()
	//log.SetGlobalPlatformName("Finvest")
	//log.SetGlobalTimestampFieldName("timestamp")
	//log.SetGlobalLogLevel(log.WarnLevel)
	//log.Warn().Msg("Hello test test")

	log.Panic().Msg("sdsdsd")
}
