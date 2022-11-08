package stdlog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type optionManagerTestSuite struct {
	suite.Suite
}

func TestOptionManagerTestSuite(t *testing.T) {
	suite.Run(t, new(optionManagerTestSuite))
}

func (s *optionManagerTestSuite) TestNewOptionManager() {
	optManager := &OptionManager{
		JSONFieldFormatter: &JSONFieldFormatter{
			PlatformNameKey: PLATFORM_NAME_KEY,
		},
	}
	assert.NotNil(s.T(), optManager)
}

func (s *optionManagerTestSuite) TestDisabledPlatformNameKey() {
	var buf bytes.Buffer

	lg := NewLogger(&buf, &OptionManager{
		DisabledPlatformNameKey: true,
		PlatformName:            "Finvest",
		JSONFieldFormatter:      &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "platform_name", val: "Finvest"},
	}

	// Check key exist
	for _, want := range wants {
		assert.Equal(s.T(), false, checkKeyMap(result, want.key))
	}

	// Check expected value
	for _, want := range wants {
		assert.Equal(s.T(), false, checkValueMap(result, want.key, want.val))
	}
}

func (s *optionManagerTestSuite) TestGetPlatformNameKey() {
	const givenPlatformNameVal = "Finvest"
	optManager := &OptionManager{
		PlatformName:       givenPlatformNameVal,
		JSONFieldFormatter: &JSONFieldFormatter{},
	}
	assert.Equal(s.T(), givenPlatformNameVal, optManager.PlatformName)
}

type jsonFieldFormatterTestSuite struct {
	suite.Suite
}

func TestJSONFieldFormatterTestSuite(t *testing.T) {
	suite.Run(t, new(jsonFieldFormatterTestSuite))
}

func (s *optionManagerTestSuite) TestNewJSONFieldFormatter() {
	jsonFieldFormatter := &JSONFieldFormatter{}
	assert.NotNil(s.T(), jsonFieldFormatter)
}

func (s *optionManagerTestSuite) TestNewJSONFieldFormatterWithPlatformNameDefault() {
	var buf bytes.Buffer
	const givenPlatformNameKey = "platform_name"

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "INFO"},
		{key: "msg", val: "test"},
		{key: givenPlatformNameKey, val: "Finvest"},
	}

	// Check key exist
	for _, want := range wants {
		assert.Equal(s.T(), true, checkKeyMap(result, want.key))
	}

	// Check expected value
	for _, want := range wants {
		assert.Equal(s.T(), true, checkValueMap(result, want.key, want.val))
	}
}

func (s *optionManagerTestSuite) TestNewJSONFieldFormatterWithPlatformNameNotDefault() {
	var buf bytes.Buffer
	const givenPlatformNameKey = "app_name"

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{PlatformNameKey: givenPlatformNameKey},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "INFO"},
		{key: "msg", val: "test"},
		{key: givenPlatformNameKey, val: "Finvest"},
	}

	// Check key exist
	for _, want := range wants {
		assert.Equal(s.T(), true, checkKeyMap(result, want.key))
	}

	// Check expected value
	for _, want := range wants {
		assert.Equal(s.T(), true, checkValueMap(result, want.key, want.val))
	}
}
