package stdlog

import (
	"fmt"

	"golang.org/x/exp/slog"
)

type Level int

const (
	DEBUG_LEVEL  Level = -4
	INFO_LEVEL   Level = 0
	WARN_LEVEL   Level = 4
	ERROR_LEVEL  Level = 8
	SILENT_LEVEL Level = 12
)

func (l Level) String() string {
	str := func(base string, val Level) string {
		if val == 0 {
			return base
		}
		return fmt.Sprintf("%s%+d", base, val)
	}

	switch {
	case l < INFO_LEVEL:
		return str("DEBUG", l-DEBUG_LEVEL)
	case l < WARN_LEVEL:
		return str("INFO", l)
	case l < ERROR_LEVEL:
		return str("WARN", l-WARN_LEVEL)
	case l == SILENT_LEVEL:
		return str("SILENT", l-SILENT_LEVEL)
	default:
		return str("ERROR", l-ERROR_LEVEL)
	}
}

func (l Level) Level() slog.Level {
	return slog.Level(l)
}
