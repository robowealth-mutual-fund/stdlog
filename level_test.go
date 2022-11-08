package stdlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slog"
)

type levelTestSuite struct {
	suite.Suite
}

func TestLevelTestSuite(t *testing.T) {
	suite.Run(t, new(levelTestSuite))
}

func (s *levelTestSuite) TestLevelDebugString() {
	assert.Equal(s.T(), "DEBUG", DebugLevel.String())
}

func (s *levelTestSuite) TestLevelInfoString() {
	assert.Equal(s.T(), "INFO", InfoLevel.String())
}

func (s *levelTestSuite) TestLevelWarnString() {
	assert.Equal(s.T(), "WARN", WarnLevel.String())
}

func (s *levelTestSuite) TestLevelErrorString() {
	assert.Equal(s.T(), "ERROR", ErrorLevel.String())
}

func (s *levelTestSuite) TestLevelSilentString() {
	assert.Equal(s.T(), "SILENT", SilentLevel.String())
}

func (s *levelTestSuite) TestGetDebugLevel() {
	assert.Equal(s.T(), slog.Level(DebugLevel), DebugLevel.Level())
}

func (s *levelTestSuite) TestGetInfoLevel() {
	assert.Equal(s.T(), slog.Level(InfoLevel), InfoLevel.Level())
}

func (s *levelTestSuite) TestGetWarnLevel() {
	assert.Equal(s.T(), slog.Level(WarnLevel), WarnLevel.Level())
}

func (s *levelTestSuite) TestGetErrorLevel() {
	assert.Equal(s.T(), slog.Level(ErrorLevel), ErrorLevel.Level())
}

func (s *levelTestSuite) TestGetSilentLevel() {
	assert.Equal(s.T(), slog.Level(SilentLevel), SilentLevel.Level())
}

func (s *levelTestSuite) TestStringNotReturnBase() {
	const TestLevel Level = 99
	assert.Equal(s.T(), "ERROR+91", TestLevel.String())
}
