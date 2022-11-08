package stdlog

import (
	"io"
	"os"

	"golang.org/x/exp/slog"
)

type GlobalSetting struct {
	writer       io.Writer
	level        Level
	platformName string
}

func NewGlobalSetting() *GlobalSetting {
	return &GlobalSetting{}
}

func (gs *GlobalSetting) WithWriter(w io.Writer) *GlobalSetting {
	gs.writer = w
	return gs
}

func (gs *GlobalSetting) WithLevel(level Level) *GlobalSetting {
	gs.level = level
	return gs
}

func (gs *GlobalSetting) WithPlatformName(name string) *GlobalSetting {
	gs.platformName = name
	return gs
}

func (gs *GlobalSetting) Configure() {
	var attrs []slog.Attr

	if gs.platformName != "" {
		attrs = append(attrs, slog.Any(platformName, gs.platformName))
	}

	logLevel = gs.level

	if gs.writer == nil {
		gs.writer = os.Stdout
	}

	h := slog.HandlerOptions{Level: gs.level}.NewJSONHandler(gs.writer).WithAttrs(attrs)
	log = slog.New(h)
}
