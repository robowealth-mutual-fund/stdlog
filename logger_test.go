package stdlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type loggerTestSuite struct {
	suite.Suite
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(loggerTestSuite))
}

func (s *loggerTestSuite) TestNewLogger() {
	lg := NewLogger(os.Stdout, &OptionManager{
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	assert.NotNil(s.T(), lg)
}

func (s *loggerTestSuite) TestNewLogDebugLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
		{key: "msg", val: "test"},
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

func (s *loggerTestSuite) TestNewLogInfoLevel() {
	var buf bytes.Buffer

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

func (s *loggerTestSuite) TestNewLogWarnLevel() {
	var buf bytes.Buffer

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
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
		{key: "level", val: "WARN"},
		{key: "msg", val: "test"},
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

func (s *loggerTestSuite) TestNewLogErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
		{key: "msg", val: "test"},
		{key: "platform_name", val: "Finvest"},
		{key: "err", val: givenErr.Error()},
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

func (s *loggerTestSuite) TestNewLogAttrsDebugLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(DEBUG_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
		{key: "msg", val: "test"},
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

func (s *loggerTestSuite) TestNewLogAttrsInfoLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(INFO_LEVEL, "test", givenAttrs)

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

func (s *loggerTestSuite) TestNewLogAttrsWarnLevel() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_1": "attr value 1"}

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(WARN_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "WARN"},
		{key: "msg", val: "test"},
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

func (s *loggerTestSuite) TestNewLogAttrsErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("attr error value 1")
	givenAttrs := Attrs{"attr_err_1": givenErr}

	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(ERROR_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
		{key: "msg", val: "test"},
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

func (s *loggerTestSuite) TestLogDebugLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogInfoLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogWarnLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogAttrsDebugLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(DEBUG_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogAttrsInfoLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(INFO_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogAttrsWarnLevelWithSilent() {
	var buf bytes.Buffer
	givenAttrs := Attrs{"attr_key_1": "attr value 1"}

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(WARN_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogAttrsErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")
	givenAttrs := Attrs{"attr_err_key_1": givenErr}

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.LogAttrs(ERROR_LEVEL, "test", givenAttrs)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogDebugLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogDebugLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogDebugLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogDebugLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogInfoLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogInfoLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogInfoLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogInfoLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogWarnLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogWarnLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogWarnLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogWarnLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *loggerTestSuite) TestLogErrorLevelWithDebug() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogErrorLevelWithInfo() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogErrorLevelWithWarn() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *loggerTestSuite) TestLogErrorLevelWithError() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	lg := NewLogger(&buf, &OptionManager{
		PlatformName:       "Finvest",
		JSONFieldFormatter: &JSONFieldFormatter{},
	})
	lg.Error("test", givenErr)

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
