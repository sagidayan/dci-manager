package v1

import (
	"github.com/rh-ecosystem-edge/dci-manager/config/common"
)

type singleSuite struct {
	agentName     string
	componentName string
	targetVersion string
	appVersion    string
	name          string
	steps         []common.TestStep
	setup         []common.TestStep
	variables     common.TestVariables
}

var _ common.Test = &singleSuite{}

func (s *singleSuite) Name() string {
	return s.name
}

func (s *singleSuite) AppVersion() string {
	return s.appVersion
}

func (s *singleSuite) TargetVersion() string {
	return s.targetVersion
}

func (s *singleSuite) AgentName() string {
	return s.agentName
}

func (s *singleSuite) ComponentName() string {
	return s.componentName
}

func (s *singleSuite) Equals(c common.Test) bool {
	return s.agentName == c.AgentName() &&
		s.componentName == c.ComponentName() &&
		s.targetVersion == c.TargetVersion() &&
		s.appVersion == c.AppVersion() &&
		s.name == c.Name() &&
		common.EqualSteps(s.steps, c.Steps()) &&
		common.EqualSteps(s.setup, c.Setup()) &&
		common.EqualVariables(s.Variables(), c.Variables())
}

func (s *singleSuite) Modified(c common.Test) bool {
	return s.agentName == c.AgentName() &&
		s.componentName == c.ComponentName() &&
		s.targetVersion == c.TargetVersion() &&
		s.appVersion == c.AppVersion() &&
		s.name == c.Name() && (!common.EqualSteps(s.steps, c.Steps()) ||
		!common.EqualVariables(s.variables, c.Variables()))
}

func (s *singleSuite) Steps() []common.TestStep {
	return s.steps
}

func (s *singleSuite) Setup() []common.TestStep {
	return s.setup
}

func (s *singleSuite) Variables() common.TestVariables {
	return s.variables
}
