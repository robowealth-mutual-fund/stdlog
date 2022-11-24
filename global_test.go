package stdlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type globalLogSettingTestSuite struct {
	suite.Suite
}

func TestGlobalLogSettingTestSuite(t *testing.T) {
	suite.Run(t, new(globalLogSettingTestSuite))
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithDebugLevel() {
	assert.Equal(s.T(), DEBUG_LEVEL, GlobalSetting{Level: DEBUG_LEVEL}.Level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithInfoLevel() {
	assert.Equal(s.T(), INFO_LEVEL, GlobalSetting{Level: INFO_LEVEL}.Level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithWarnLevel() {
	assert.Equal(s.T(), WARN_LEVEL, GlobalSetting{Level: WARN_LEVEL}.Level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithErrorLevel() {
	assert.Equal(s.T(), ERROR_LEVEL, GlobalSetting{Level: ERROR_LEVEL}.Level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithSilentLevel() {
	assert.Equal(s.T(), SILENT_LEVEL, GlobalSetting{Level: SILENT_LEVEL}.Level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithPlatformName() {
	assert.Equal(s.T(), "Test", GlobalSetting{PlatformName: "Test"}.PlatformName)
}

func (s *globalLogSettingTestSuite) TestConfigureWithPlatformNameWithDefaultLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, PlatformName: "Test"}
	globalSetting.Configure()

	log.Info("test")
	fmt.Println(buf.String())
	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "INFO"},
		{key: "platform_name", val: "Test"},
	}

	// Check key exist
	for _, want := range wants {
		assert.Equal(s.T(), true, checkKeyMap(result, want.key))
	}

	// Check expected value
	for _, want := range wants {
		assert.Equal(s.T(), true, checkValueMap(result, want.key, want.val))
	}

	assert.Equal(s.T(), INFO_LEVEL, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithDebugLevel() {
	globalSetting := &GlobalSetting{Level: DEBUG_LEVEL}
	globalSetting.Configure()
	assert.Equal(s.T(), DEBUG_LEVEL, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithInfoLevel() {
	globalSetting := &GlobalSetting{Level: INFO_LEVEL}
	globalSetting.Configure()
	assert.Equal(s.T(), INFO_LEVEL, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithWarnLevel() {
	globalSetting := &GlobalSetting{Level: WARN_LEVEL}
	globalSetting.Configure()
	assert.Equal(s.T(), WARN_LEVEL, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithErrorLevel() {
	globalSetting := &GlobalSetting{Level: ERROR_LEVEL}
	globalSetting.Configure()
	assert.Equal(s.T(), ERROR_LEVEL, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithSilentLevel() {
	globalSetting := &GlobalSetting{Level: SILENT_LEVEL}
	globalSetting.Configure()
	assert.Equal(s.T(), SILENT_LEVEL, logLevel)
}
