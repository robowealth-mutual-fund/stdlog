package stdlog

const (
	PLATFORM_NAME_KEY = "platform_name"
)

type OptionManager struct {
	DisabledPlatformNameKey bool
	PlatformName            string
	JSONFieldFormatter      *JSONFieldFormatter
}

func (om *OptionManager) WithJSONFieldFormatter(jsonFieldFormatter *JSONFieldFormatter) *OptionManager {
	om.JSONFieldFormatter = jsonFieldFormatter
	return om
}

type JSONFieldFormatter struct {
	PlatformNameKey string
}

func (jff *JSONFieldFormatter) PlatformNameKeyWithDefault() string {
	if jff.PlatformNameKey != "" {
		return jff.PlatformNameKey
	}
	return PLATFORM_NAME_KEY
}
