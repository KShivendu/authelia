package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OIDCSuite struct {
	*RodSuite
}

func NewOIDCSuite() *OIDCSuite {
	return &OIDCSuite{
		RodSuite: &RodSuite{
			Name: oidcSuiteName,
		},
	}
}

func (s *OIDCSuite) TestOIDCScenario() {
	suite.Run(s.T(), NewOIDCScenario())
}

func (s *OIDCSuite) SetupSuite() {
	s.LoadEnvironment()
}

func TestOIDCSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewOIDCSuite())
}
