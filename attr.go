package stdlog

import (
	"golang.org/x/exp/slog"
)

//type Attr struct {
//	key   string
//	value value
//}
//
//func AnyValue(key string, val any) Attr {
//	return Attr{
//		key: key,
//		value: value{
//			num: 0,
//			any: val,
//		},
//	}
//}
//
//type Attrs []Attr
//
//func (a Attrs) Convert() []slog.Attr {
//	attrs := make([]slog.Attr, len(a))
//	for i, attr := range a {
//		attrs[i] = slog.Attr{
//			Key:   attr.key,
//			Value: slog.AnyValue(attr.value.any),
//		}
//	}
//	return attrs
//}

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
