package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HAProxySuite struct {
	*RodSuite
}

func NewHAProxySuite() *HAProxySuite {
	return &HAProxySuite{
		RodSuite: &RodSuite{
			Name: haproxySuiteName,
		},
	}
}

func (s *HAProxySuite) Test1FAScenario() {
	suite.Run(s.T(), New1FAScenario())
}

func (s *HAProxySuite) Test2FAScenario() {
	suite.Run(s.T(), New2FAScenario())
}

func (s *HAProxySuite) TestCustomHeaders() {
	suite.Run(s.T(), NewCustomHeadersScenario())
}

func (s *HAProxySuite) SetupSuite() {
	s.LoadEnvironment()
}

func TestHAProxySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewHAProxySuite())
}
