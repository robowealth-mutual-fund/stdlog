package stdlog

import "github.com/robowealth-mutual-fund/stdlog/internal/constants"

func SetGlobalApplicationName(name string) {
	log = log.With(constants.APPLICATION_NAME_KEY, name)
}

func SetGlobalLogLevel(level Level) {
	logLevel = level
	levelVar.Set(logLevel.Level())
}

func SetEnabledReportCaller() {
	reportCaller = true
}

func ReportCaller() bool {
	return reportCaller
}
