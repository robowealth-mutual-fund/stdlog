package stdlog

import (
	"bytes"
	"encoding/json"
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
	assert.Equal(s.T(), DebugLevel, NewGlobalSetting().WithLevel(DebugLevel).level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithInfoLevel() {
	assert.Equal(s.T(), InfoLevel, NewGlobalSetting().WithLevel(InfoLevel).level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithWarnLevel() {
	assert.Equal(s.T(), WarnLevel, NewGlobalSetting().WithLevel(WarnLevel).level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithErrorLevel() {
	assert.Equal(s.T(), ErrorLevel, NewGlobalSetting().WithLevel(ErrorLevel).level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithSilentLevel() {
	assert.Equal(s.T(), SilentLevel, NewGlobalSetting().WithLevel(SilentLevel).level)
}

func (s *globalLogSettingTestSuite) TestNewGlobalLogSettingWithPlatformName() {
	assert.Equal(s.T(), "Test", NewGlobalSetting().WithPlatformName("Test").platformName)
}

func (s *globalLogSettingTestSuite) TestConfigureWithPlatformNameWithDefaultLevel() {
	var buf bytes.Buffer

	NewGlobalSetting().WithWriter(&buf).WithPlatformName("Test").Configure()
	log.Info("test")

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

	assert.Equal(s.T(), InfoLevel, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithDebugLevel() {
	NewGlobalSetting().WithLevel(DebugLevel).Configure()
	assert.Equal(s.T(), DebugLevel, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithInfoLevel() {
	NewGlobalSetting().WithLevel(InfoLevel).Configure()
	assert.Equal(s.T(), InfoLevel, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithWarnLevel() {
	NewGlobalSetting().WithLevel(WarnLevel).Configure()
	assert.Equal(s.T(), WarnLevel, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithErrorLevel() {
	NewGlobalSetting().WithLevel(ErrorLevel).Configure()
	assert.Equal(s.T(), ErrorLevel, logLevel)
}

func (s *globalLogSettingTestSuite) TestConfigureWithSilentLevel() {
	NewGlobalSetting().WithLevel(SilentLevel).Configure()
	assert.Equal(s.T(), SilentLevel, logLevel)
}
