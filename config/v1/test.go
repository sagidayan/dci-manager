package v1

import "github.com/rh-ecosystem-edge/dci-manager/config/common"

type Test struct {
	Name           string    `yaml:"name" json:"name"`
	TargetVersions []string  `yaml:"target_versions" json:"target_versions"`
	AppVersions    []string  `yaml:"app_versions" json:"app_versions"`
	Steps          []Step    `yaml:"steps,omitempty" json:"steps,omitempty"`
	Setup          []Step    `yaml:"setup,omitempty" json:"setup,omitempty"`
	Variables      variables `yaml:"variables,omitempty" json:"variables,omitempty"`
}

func (t *Test) GetSteps() []common.TestStep {
	var resp []common.TestStep = []common.TestStep{}
	for _, s := range t.Steps {
		st := s
		resp = append(resp, &st)
	}
	return resp
}

func (t *Test) GetSetup() []common.TestStep {
	var resp []common.TestStep = []common.TestStep{}
	for _, s := range t.Setup {
		st := s
		resp = append(resp, &st)
	}
	return resp
}

func (t *Test) GetVariables() common.TestVariables {
	return &t.Variables
}
