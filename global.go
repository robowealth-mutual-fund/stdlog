package stdlog

import (
	"io"
	"os"

	"golang.org/x/exp/slog"
)

type GlobalSetting struct {
	Writer       io.Writer
	Level        Level
	PlatformName string
}

func (gs *GlobalSetting) Configure() {
	var attrs []slog.Attr

	if gs.PlatformName != "" {
		attrs = append(attrs, slog.Any(PLATFORM_NAME_KEY, gs.PlatformName))
	}

	logLevel = gs.Level

	if gs.Writer == nil {
		gs.Writer = os.Stdout
	}

	h := slog.HandlerOptions{Level: gs.Level}.NewJSONHandler(gs.Writer).WithAttrs(attrs)
	log = slog.New(h)
}
