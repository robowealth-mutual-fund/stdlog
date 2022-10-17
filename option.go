package stdlog

type logOptions struct {
	level           Level
	platformName    string
	enableTimestamp bool

	// Field name
	timestampFieldName string
	levelFieldName     string
	messageFieldName   string

	timeFieldFormat TimeFormat
}

type Option struct {
	fn func(*logOptions)
}

func (opt *Option) Apply(do *logOptions) {
	opt.fn(do)
}

func newLogOption(fn func(*logOptions)) *Option {
	return &Option{
		fn: fn,
	}
}

func WithGlobalLevel(level Level) *Option {
	return newLogOption(func(o *logOptions) {
		o.level = level
	})
}

func WithPlatformName(name string) *Option {
	return newLogOption(func(o *logOptions) {
		o.platformName = name
	})
}

func WithCustomTimestampFieldName(key string) *Option {
	return newLogOption(func(o *logOptions) {
		o.timestampFieldName = key
	})
}

func WithTimestamp(isEnabled bool) *Option {
	return newLogOption(func(o *logOptions) {
		o.enableTimestamp = isEnabled
	})
}

func WithCustomTimeFieldFormat() {

}
