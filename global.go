package stdlog

import "github.com/robowealth-mutual-fund/stdlog/internal/constants"

func SetGlobalPlatformName(name string) {
	log = log.With(constants.PLATFORM_NAME_KEY, name)
}

func SetGlobalLogLevel(level Level) {
	logLevel = level
	levelVar.Set(logLevel.Level())
}

func SetEnabledReportCaller() {
	reportCaller = true
}
