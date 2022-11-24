package stdlog

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type logTestSuite struct {
	suite.Suite
}

func TestLogTestSuite(t *testing.T) {
	suite.Run(t, new(logTestSuite))
}

func (s *logTestSuite) TestLogDebugLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
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
}

func (s *logTestSuite) TestLogInfoLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	Info("test")

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
}

func (s *logTestSuite) TestLogWarnLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "WARN"},
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
}

func (s *logTestSuite) TestLogErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
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
}

func (s *logTestSuite) TestLogAttrDebugLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	LogAttrs(DEBUG_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "DEBUG"},
		{key: "platform_name", val: "Test"},
		{key: "attr_key_1", val: "attr value 1"},
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

func (s *logTestSuite) TestLogAttrInfoLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	LogAttrs(INFO_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

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
		{key: "attr_key_1", val: "attr value 1"},
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

func (s *logTestSuite) TestLogAttrWarnLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	LogAttrs(WARN_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "WARN"},
		{key: "platform_name", val: "Test"},
		{key: "attr_key_1", val: "attr value 1"},
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

func (s *logTestSuite) TestLogAttrErrorLevel() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL, PlatformName: "Test"}
	globalSetting.Configure()
	LogAttrs(ERROR_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.T().Fatal(err)
	}

	wants := []struct {
		key string
		val string
	}{
		{key: "level", val: "ERROR"},
		{key: "platform_name", val: "Test"},
		{key: "attr_key_1", val: "attr value 1"},
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

func (s *logTestSuite) TestLogDebugLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogInfoLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogWarnLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogAttrsDebugLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	LogAttrs(DEBUG_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogAttrsInfoLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	LogAttrs(INFO_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogAttrsWarnLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	LogAttrs(WARN_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogAttrsErrorLevelWithSilent() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: SILENT_LEVEL}
	globalSetting.Configure()
	LogAttrs(ERROR_LEVEL, "test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogDebugLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogDebugLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogDebugLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogDebugLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogInfoLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogInfoLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogInfoLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogInfoLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogWarnLevelWithDebug() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogWarnLevelWithInfo() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogWarnLevelWithWarn() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogWarnLevelWithError() {
	var buf bytes.Buffer

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *logTestSuite) TestLogErrorLevelWithDebug() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: DEBUG_LEVEL}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogErrorLevelWithInfo() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: INFO_LEVEL}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogErrorLevelWithWarn() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: WARN_LEVEL}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}

func (s *logTestSuite) TestLogErrorLevelWithError() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	globalSetting := &GlobalSetting{Writer: &buf, Level: ERROR_LEVEL}
	globalSetting.Configure()
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.NoError(s.T(), err)
	}
}
