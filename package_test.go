package stdlog

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type packageTestSuite struct {
	suite.Suite
}

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(packageTestSuite))
}
