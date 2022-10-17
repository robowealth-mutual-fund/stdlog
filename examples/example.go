package main

import (
	log "github.com/robowealth-mutual-fund/stdlog"
)

func main() {
	//fluentLog := log.New(os.Stderr,
	//	log.WithPlatformName("Odini"),
	//	log.WithGlobalLevel(log.WarnLevel),
	//	log.WithTimestamp(true),
	//)

	//fluentLog.Warn().Str("some_key", "Hello World").Msg("This is standard log lib!")
	log.SetTimestamp()
	log.Warn().Msg("sdsdsd")
}
