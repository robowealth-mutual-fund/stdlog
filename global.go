package stdlog

import (
	"os"

	"golang.org/x/exp/slog"
)

type GlobalSetting struct {
	level        Level
	platformName string
}

func NewGlobalSetting() *GlobalSetting {
	return &GlobalSetting{}
}

func (gs *GlobalSetting) WithLevel(level Level) *GlobalSetting {
	gs.level = level
	logLevel = level
	return gs
}

func (gs *GlobalSetting) WithPlatformName(name string) *GlobalSetting {
	gs.platformName = name
	return gs
}

func (gs *GlobalSetting) Build() {
	var attrs []slog.Attr

	if gs.platformName != "" {
		attrs = append(attrs, slog.Any(platformName, gs.platformName))
	}

	h := slog.HandlerOptions{Level: gs.level}.NewJSONHandler(os.Stdout).WithAttrs(attrs)
	log = slog.New(h)
}
