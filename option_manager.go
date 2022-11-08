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

func (om *OptionManager) WithDisabledPlatformNameKey() *OptionManager {
	om.disabledPlatformNameKey = true
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

func (ff *JSONFieldFormatter) PlatformNameKey() string {
	return ff.platformNameKey
}

func (ff *JSONFieldFormatter) PlatformNameKeyWithDefault() string {
	if ff.PlatformNameKey() != "" {
		return ff.platformNameKey
	}
	return platformName
}
