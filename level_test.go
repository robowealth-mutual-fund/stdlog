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
	assert.Equal(s.T(), "DEBUG", DEBUG_LEVEL.String())
}

func (s *levelTestSuite) TestLevelInfoString() {
	assert.Equal(s.T(), "INFO", INFO_LEVEL.String())
}

func (s *levelTestSuite) TestLevelWarnString() {
	assert.Equal(s.T(), "WARN", WARN_LEVEL.String())
}

func (s *levelTestSuite) TestLevelErrorString() {
	assert.Equal(s.T(), "ERROR", ERROR_LEVEL.String())
}

func (s *levelTestSuite) TestLevelSilentString() {
	assert.Equal(s.T(), "SILENT", SILENT_LEVEL.String())
}

func (s *levelTestSuite) TestGetDebugLevel() {
	assert.Equal(s.T(), slog.Level(DEBUG_LEVEL), DEBUG_LEVEL.Level())
}

func (s *levelTestSuite) TestGetInfoLevel() {
	assert.Equal(s.T(), slog.Level(INFO_LEVEL), INFO_LEVEL.Level())
}

func (s *levelTestSuite) TestGetWarnLevel() {
	assert.Equal(s.T(), slog.Level(WARN_LEVEL), WARN_LEVEL.Level())
}

func (s *levelTestSuite) TestGetErrorLevel() {
	assert.Equal(s.T(), slog.Level(ERROR_LEVEL), ERROR_LEVEL.Level())
}

func (s *levelTestSuite) TestGetSilentLevel() {
	assert.Equal(s.T(), slog.Level(SILENT_LEVEL), SILENT_LEVEL.Level())
}

func (s *levelTestSuite) TestStringNotReturnBase() {
	const TestLevel Level = 99
	assert.Equal(s.T(), "ERROR+91", TestLevel.String())
}
