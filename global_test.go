package stdlog

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/suite"
	"os"
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
	logger.Info().Msg("test platform name")
	reader := bufio.NewReader(os.Stdout)
	text, _ := reader.ReadString('\n')
	fmt.Println("dssdds", text)
}

func (s *logTestSuite) SetGlobalTimestampFieldName() {
	SetGlobalPlatformName("test plat form")
}
