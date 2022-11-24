package stdlog

import (
	"golang.org/x/exp/slog"
)

// An Attrs is a key-value pair.
type Attrs map[string]interface{}

func (a Attrs) convert() []slog.Attr {
	var attrs []slog.Attr
	for s, i := range a {
		attrs = append(attrs, slog.Attr{
			Key:   s,
			Value: slog.AnyValue(i),
		})
	}
	return attrs
}
