package stdlog

const (
	platformName = "platform_name"
)

type OptionManager struct {
	disabledPlatformNameKey bool
	platformName            string
	jsonFieldManager        *JSONFieldFormatter
}

func NewOptionManager(jsonfieldManager *JSONFieldFormatter) *OptionManager {
	return &OptionManager{
		jsonFieldManager: jsonfieldManager,
	}
}

func NewOptionManagerWithDefaultConfig() *OptionManager {
	jsonFieldManager := NewJSONFieldFormatter().WithPlatformNameKey("platform_name")

	return &OptionManager{
		jsonFieldManager: jsonFieldManager,
	}
}

func (om *OptionManager) WithDisabledPlatformNameKey(disabled bool) *OptionManager {
	om.disabledPlatformNameKey = disabled
	return om
}

func (om *OptionManager) DisabledPlatformNameKey() bool {
	return om.disabledPlatformNameKey
}

func (om *OptionManager) WithPlatformName(name string) *OptionManager {
	om.platformName = name
	return om
}

func (om *OptionManager) PlatformName() string {
	return om.platformName
}

type JSONFieldFormatter struct {
	platformNameKey  string
	timestampNameKey string
}

func NewJSONFieldFormatter() *JSONFieldFormatter {
	return &JSONFieldFormatter{}
}

func (ff *JSONFieldFormatter) WithPlatformNameKey(name string) *JSONFieldFormatter {
	ff.platformNameKey = name
	return ff
}

func (ff *JSONFieldFormatter) PlatformNameKeyWithDefault() string {
	if ff.platformNameKey != "" {
		return ff.platformNameKey
	}
	return platformName
}
