package stdlog

import (
	"fmt"

	"golang.org/x/exp/slog"
)

type Level int

const (
	DebugLevel  Level = -4
	InfoLevel   Level = 0
	WarnLevel   Level = 4
	ErrorLevel  Level = 8
	SilentLevel Level = 12
)

func (l Level) String() string {
	str := func(base string, val Level) string {
		if val == 0 {
			return base
		}
		return fmt.Sprintf("%s%+d", base, val)
	}

	switch {
	case l < InfoLevel:
		return str("DEBUG", l-DebugLevel)
	case l < WarnLevel:
		return str("INFO", l)
	case l < ErrorLevel:
		return str("WARN", l-WarnLevel)
	case l == SilentLevel:
		return str("SILENT", l-SilentLevel)
	default:
		return str("ERROR", l-ErrorLevel)
	}
}

func (l Level) Level() slog.Level {
	return slog.Level(l)
}
