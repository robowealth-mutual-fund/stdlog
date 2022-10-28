package stdlog

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	errCode "github.com/robowealth-mutual-fund/stdlog/constant/errors"
	"github.com/robowealth-mutual-fund/stdlog/constant/errors/otp"
)

type eventTestSuite struct {
	suite.Suite
}

func TestEventTestSuite(t *testing.T) {
	suite.Run(t, new(eventTestSuite))
}

func (s *eventTestSuite) TestEventMsgWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Msg("test")
}

func (s *eventTestSuite) TestEventMsgfWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Msgf("%s", "test")
}

func (s *eventTestSuite) TestEventInterfaceWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Interface("test_key", "test").Msg("test info")
}

func (s *eventTestSuite) TestEventErrWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Err(otp.INCORRECT).Msg("test info")
}

func (s *eventTestSuite) TestEventErrsWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Errs("test_errors", []errCode.Interface{otp.INCORRECT}).Msg("test info")
}

func (s *eventTestSuite) TestEventStrWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Str("test", "test log").Msg("test info")
}

func (s *eventTestSuite) TestEventStrsWithLogInfo() {
	testLogger := New(os.Stdout)
	testLogger.Info().Strs("tests", []string{"test log 1", "test log 2"}).Msg("test info")
}

func (s *eventTestSuite) TestEventMsgWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Msg("test debug")
}

func (s *eventTestSuite) TestEventMsgfWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Msgf("%s", "test debug")
}

func (s *eventTestSuite) TestEventInterfaceWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Interface("test_key", "test").Msg("test debug")
}

func (s *eventTestSuite) TestEventErrWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Err(otp.INCORRECT).Msg("test debug")
}

func (s *eventTestSuite) TestEventErrsWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Errs("test_errors", []errCode.Interface{otp.INCORRECT}).Msg("test debug")
}

func (s *eventTestSuite) TestEventStrWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Str("test", "test log").Msg("test debug")
}

func (s *eventTestSuite) TestEventStrsWithLogDebug() {
	testLogger := New(os.Stdout)
	testLogger.Debug().Strs("tests", []string{"test log 1", "test log 2"}).Msg("test debug")
}

func (s *eventTestSuite) TestEventMsgWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Msg("test warn")
}

func (s *eventTestSuite) TestEventMsgfWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Msgf("%s", "test warn")
}

func (s *eventTestSuite) TestEventInterfaceWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Interface("test_key", "test").Msg("test warn")
}

func (s *eventTestSuite) TestEventErrWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Err(otp.INCORRECT).Msg("test warn")
}

func (s *eventTestSuite) TestEventErrsWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Errs("test_errors", []errCode.Interface{otp.INCORRECT}).Msg("test warn")
}

func (s *eventTestSuite) TestEventStrWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Str("test", "test log").Msg("test warn")
}

func (s *eventTestSuite) TestEventStrsWithLogWarn() {
	testLogger := New(os.Stdout)
	testLogger.Warn().Strs("tests", []string{"test log 1", "test log 2"}).Msg("test warn")
}

func (s *eventTestSuite) TestEventMsgWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Msg("test error")
}

func (s *eventTestSuite) TestEventMsgfWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Msgf("%s", "test error")
}

func (s *eventTestSuite) TestEventInterfaceWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Interface("test_key", "test").Msg("test error")
}

func (s *eventTestSuite) TestEventErrWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Err(otp.INCORRECT).Msg("test error")
}

func (s *eventTestSuite) TestEventErrsWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Errs("test_errors", []errCode.Interface{otp.INCORRECT}).Msg("test error")
}

func (s *eventTestSuite) TestEventStrWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Str("test", "test log").Msg("test error")
}

func (s *eventTestSuite) TestEventStrsWithLogError() {
	testLogger := New(os.Stdout)
	testLogger.Error().Strs("tests", []string{"test log 1", "test log 2"}).Msg("test error")
}

func (s *eventTestSuite) TestLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal()
}

func (s *eventTestSuite) TestEventInterfaceWithLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal().Interface("test_key", "test")
}

func (s *eventTestSuite) TestEventErrWithLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal().Err(otp.INCORRECT)
}

func (s *eventTestSuite) TestEventErrsWithLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal().Errs("test_errors", []errCode.Interface{otp.INCORRECT})
}

func (s *eventTestSuite) TestEventStrWithLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal().Str("test", "test log")
}

func (s *eventTestSuite) TestEventStrsWithLogFatal() {
	testLogger := New(os.Stdout)
	testLogger.Fatal().Strs("tests", []string{"test log 1", "test log 2"})
}

func (s *eventTestSuite) TestEventMsgWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Msg("test")
	})
}

func (s *eventTestSuite) TestEventMsgfWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Msgf("%s", "test")
	})
}

func (s *eventTestSuite) TestEventInterfaceWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Interface("test_key", "test").Msg("test panic")
	})
}

func (s *eventTestSuite) TestEventErrWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Err(otp.INCORRECT).Msg("test panic")
	})
}

func (s *eventTestSuite) TestEventErrsWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Errs("test_errors", []errCode.Interface{otp.INCORRECT}).Msg("test panic")
	})
}

func (s *eventTestSuite) TestEventStrWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Str("test", "test log").Msg("test panic")
	})
}

func (s *eventTestSuite) TestEventStrsWithLogPanic() {
	testLogger := New(os.Stdout)
	assert.Panics(s.T(), func() {
		testLogger.Panic().Strs("some_key", []string{"test panic"}).Msg("test panic")
	})
}

func (s *eventTestSuite) TestWith() {
	testLogger := New(os.Stdout)
	testLogger.With()
}
