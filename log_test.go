package stdlog

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
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

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
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

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
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

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevel() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
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

func (s *packageTestSuite) TestGlobalInstanceLogAttrDebugLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
	DebugWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

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

func (s *packageTestSuite) TestGlobalInstanceLogAttrInfoLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
	InfoWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

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

func (s *packageTestSuite) TestGlobalInstanceLogAttrWarnLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
	WarnWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

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

func (s *packageTestSuite) TestGlobalInstanceLogAttrErrorLevel() {
	var buf bytes.Buffer

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	SetGlobalPlatformName("Test")
	ErrorWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

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

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevelWithSilent() {
	var buf bytes.Buffer

	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevelWithSilent() {
	var buf bytes.Buffer

	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevelWithSilent() {
	var buf bytes.Buffer

	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "no output log")
	}
}

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevelWithSilent() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(SILENT_LEVEL)
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		assert.Error(s.T(), err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogAttrsDebugLevelWithSilent() {
	var buf bytes.Buffer

	SetGlobalLogLevel(SILENT_LEVEL)
	DebugWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogAttrsInfoLevelWithSilent() {
	var buf bytes.Buffer

	SetGlobalLogLevel(SILENT_LEVEL)
	InfoWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogAttrsWarnLevelWithSilent() {
	var buf bytes.Buffer

	SetGlobalLogLevel(SILENT_LEVEL)
	WarnWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogAttrsErrorLevelWithSilent() {
	var buf bytes.Buffer

	SetGlobalLogLevel(SILENT_LEVEL)
	ErrorWithAttrs("test", Attrs{"attr_key_1": "attr value 1"})

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevelWithDebug() {
	var buf bytes.Buffer

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevelWithInfo() {
	var buf bytes.Buffer

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevelWithWarn() {
	var buf bytes.Buffer

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogDebugLevelWithError() {
	var buf bytes.Buffer

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	Debug("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevelWithDebug() {
	var buf bytes.Buffer

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevelWithInfo() {
	var buf bytes.Buffer

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevelWithWarn() {
	var buf bytes.Buffer

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogInfoLevelWithError() {
	var buf bytes.Buffer

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	Info("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevelWithDebug() {
	var buf bytes.Buffer

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevelWithInfo() {
	var buf bytes.Buffer

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevelWithWarn() {
	var buf bytes.Buffer

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogWarnLevelWithError() {
	var buf bytes.Buffer

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	Warn("test")

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.Equal(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevelWithDebug() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(DEBUG_LEVEL)
	loadDefaultLoggerTest(&buf)
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevelWithInfo() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(INFO_LEVEL)
	loadDefaultLoggerTest(&buf)
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevelWithWarn() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(WARN_LEVEL)
	loadDefaultLoggerTest(&buf)
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func (s *packageTestSuite) TestGlobalInstanceLogErrorLevelWithError() {
	var buf bytes.Buffer
	givenErr := errors.New("something went wrong")

	SetGlobalLogLevel(ERROR_LEVEL)
	loadDefaultLoggerTest(&buf)
	Error("test", givenErr)

	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		s.Error(err, "error unmarshal")
	}

	s.NotEqual(0, len(result))
}

func loadDefaultLoggerTest(buf *bytes.Buffer) {
	d := &slog.HandlerOptions{
		Level:       logLevel,
		ReplaceAttr: loadDefaultReplaceAttr(),
	}

	log = slog.New(d.NewJSONHandler(buf))
}
