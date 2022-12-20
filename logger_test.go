package stdlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"

	"github.com/stretchr/testify/assert"

	"github.com/robowealth-mutual-fund/stdlog/internal/constants"
)

func (s *packageTestSuite) TestNewLogger() {
	l := NewLogger(os.Stdout, &OptionManager{
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	assert.NotNil(s.T(), l)
}

func (s *packageTestSuite) TestNewLogDebugLevel() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, DEBUG_LEVEL)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
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

func (s *packageTestSuite) TestNewLogInfoLevel() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, INFO_LEVEL)
	l.Info("test")

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
		{key: "platform_name", val: "Finvest"},
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

func (s *packageTestSuite) TestNewLogWarnLevel() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, WARN_LEVEL)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "WARN"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
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

func (s *packageTestSuite) TestNewLogErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, ERROR_LEVEL)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
		{key: constants.ERROR_KEY, val: givenErr.Error()},
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

func (s *packageTestSuite) TestNewLogAttrsDebugLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, DEBUG_LEVEL)
	l.DebugWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
		{key: "attr_1", val: "attr value 1"},
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

func (s *packageTestSuite) TestNewLogAttrsInfoLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, INFO_LEVEL)
	l.InfoWithAttrs("test", givenAttrs)

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
		{key: "platform_name", val: "Finvest"},
		{key: "attr_1", val: "attr value 1"},
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

func (s *packageTestSuite) TestNewLogAttrsWarnLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, WARN_LEVEL)
	l.WarnWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "WARN"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
		{key: "attr_1", val: "attr value 1"},
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

func (s *packageTestSuite) TestNewLogAttrsErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("attr error value 1")
	givenAttrs := Attrs{"attr_err_1": givenErr}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, ERROR_LEVEL)
	l.ErrorWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
		{key: constants.MESSAGE_KEY, val: "test"},
		{key: "platform_name", val: "Finvest"},
		{key: "attr_err_1", val: "attr error value 1"},
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

func (s *packageTestSuite) TestLogDebugLevelWithSilent() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogInfoLevelWithSilent() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogWarnLevelWithSilent() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogAttrsDebugLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.DebugWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogAttrsInfoLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.InfoWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogAttrsWarnLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.WarnWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogAttrsErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")
	givenAttrs := Attrs{"attr_err_key_1": givenErr}

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.ErrorWithAttrs("test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogDebugLevelWithDebug() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogDebugLevelWithInfo() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogDebugLevelWithWarn() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogDebugLevelWithError() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogInfoLevelWithDebug() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogInfoLevelWithInfo() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogInfoLevelWithWarn() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogInfoLevelWithError() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogWarnLevelWithDebug() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogWarnLevelWithInfo() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogWarnLevelWithWarn() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogWarnLevelWithError() {
	var buf bytes.Buffer

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestLogErrorLevelWithDebug() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogErrorLevelWithInfo() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogErrorLevelWithWarn() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *packageTestSuite) TestLogErrorLevelWithError() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	l := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	}, 0)
	l.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func checkValueMap(result map[string]any, key string, val any) bool {
	if v, ok := result[key]; ok {
		if v == val {
			return true
		}
	}
	return false
}

func checkKeyMap(val map[string]any, key string) bool {
	if _, ok := val[key]; ok {
		return true
	}
	return false
}
