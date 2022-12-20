package stdlog

import "github.com/robowealth-mutual-fund/stdlog/internal/constants"

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
	return constants.PLATFORM_NAME_KEY
}
