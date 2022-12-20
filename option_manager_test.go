package stdlog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
)

func (s *packageTestSuite) TestNewOptionManager() {
	optManager := &OptionManager{
		JSONFieldFormatter: &JSONFieldFormatter{
			PlatformNameKey: constants.PLATFORM_NAME_KEY,
		},
	}
	assert.NotNil(s.T(), optManager)
}

func (s *packageTestSuite) TestDisabledPlatformNameKey() {
	var buf bytes.Buffer

	lg := NewLogger(&buf, &OptionManager{
		DisabledPlatformNameKey: true,
		PlatformName:            "Finvest",
		JSONFieldFormatter:      &JSONFieldFormatter{},
	}, 0)
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

func (s *packageTestSuite) TestGetPlatformNameKey() {
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

func (s *packageTestSuite) TestNewJSONFieldFormatter() {
	jsonFieldFormatter := &JSONFieldFormatter{}
	assert.NotNil(s.T(), jsonFieldFormatter)
}

func (s *packageTestSuite) TestNewJSONFieldFormatterWithPlatformNameDefault() {
	var buf bytes.Buffer
	const givenPlatformNameKey = "platform_name"

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
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
		{key: constants.MESSAGE_KEY, val: "test"},
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

func (s *packageTestSuite) TestNewJSONFieldFormatterWithPlatformNameNotDefault() {
	var buf bytes.Buffer
	const givenPlatformNameKey = "app_name"

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{PlatformNameKey: givenPlatformNameKey},
	}, 0)
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
		{key: constants.MESSAGE_KEY, val: "test"},
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
