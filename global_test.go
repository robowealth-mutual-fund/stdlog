package stdlog

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type logTestSuite struct {
	suite.Suite
}

func TestLogTestSuite(t *testing.T) {
	suite.Run(t, new(logTestSuite))
}

func (s *logTestSuite) TestSetGlobalTimestamp() {
	SetGlobalTimestamp()
}

func (s *logTestSuite) TestSetGlobalLogDebugLevel() {
	SetGlobalLogLevel(DebugLevel)
}

func (s *logTestSuite) TestSetGlobalLogWarnLevel() {
	SetGlobalLogLevel(WarnLevel)
}

func (s *logTestSuite) TestSetGlobalLogErrorLevel() {
	SetGlobalLogLevel(ErrorLevel)
}

func (s *logTestSuite) TestSetGlobalLogTraceLevel() {
	SetGlobalLogLevel(TraceLevel)
}

func (s *logTestSuite) TestSetGlobalLogInfoLevel() {
	SetGlobalLogLevel(InfoLevel)
}

func (s *logTestSuite) TestSetGlobalLogFatalLevel() {
	SetGlobalLogLevel(FatalLevel)
}

func (s *logTestSuite) TestSetGlobalLogPanicLevel() {
	SetGlobalLogLevel(PanicLevel)
}

func (s *logTestSuite) TestSetGlobalLogDisabledLevel() {
	SetGlobalLogLevel(Disabled)
}

func (s *logTestSuite) TestSetGlobalLogNoLevel() {
	SetGlobalLogLevel(NoLevel)
}

func (s *logTestSuite) TestSetGlobalLogDefaultLevel() {
	SetGlobalLogLevel(99)
}

func (s *logTestSuite) TestSetGlobalLogPlatformName() {
	SetGlobalPlatformName("test platform")
}

func (s *logTestSuite) TestSetGlobalTimestampFieldName() {
	SetGlobalTimestampFieldName("test field name")
}

func (s *logTestSuite) TestSetGlobalTimeFieldFormatUnixMs() {
	SetGlobalTimeFieldFormat(TimeFormatUnixMs)
}

func (s *logTestSuite) TestSetGlobalTimeFieldFormatUnix() {
	SetGlobalTimeFieldFormat(TimeFormatUnix)
}

func (s *logTestSuite) TestSetGlobalTimeFieldFormatUnixMicro() {
	SetGlobalTimeFieldFormat(TimeFormatUnixMicro)
}

func (s *logTestSuite) TestSetGlobalTimeFieldFormatUnixNano() {
	SetGlobalTimeFieldFormat(TimeFormatUnixNano)
}

func (s *logTestSuite) TestSetGlobalTimeFieldFormatDefault() {
	SetGlobalTimeFieldFormat(99)
}

func (s *logTestSuite) TestSetGlobalLevelFieldName() {
	SetGlobalLevelFieldName("test")
}

func (s *logTestSuite) TestSetGlobalMessageFieldName() {
	SetGlobalMessageFieldName("test")
}
