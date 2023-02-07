package v1

import (
	"fmt"
	"strings"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
)

type Agent struct {
	Type          AgentType `yaml:"type" json:"type"`
	ComponentName string    `yaml:"component_name" json:"component_name"`
	Tests         []Test    `yaml:"tests" json:"tests"`
}

type AgentType string

func (a *Agent) SingleSuites() []common.Test {
	var resp []common.Test = make([]common.Test, 0)
	for _, test := range a.Tests {
		for _, targetVersion := range test.TargetVersions {
			for _, appVersion := range test.AppVersions {
				s := &singleSuite{
					name:          test.Name,
					agentName:     string(a.Type),
					componentName: a.ComponentName,
					targetVersion: targetVersion,
					appVersion:    appVersion,
					steps:         test.GetSteps(),
					setup:         test.GetSetup(),
					variables:     mergeVaribalesWithDefaults(test, a.ComponentName, targetVersion, appVersion),
				}
				resp = append(resp, s)
			}
		}
	}
	return resp

}

func mergeVaribalesWithDefaults(test Test, component string, targetVersion string, appVersion string) common.TestVariables {
	vars := variables{}
	tVars := test.GetVariables()
	for _, key := range test.GetVariables().Keys() {
		if strings.HasPrefix(key, "dcim_") {
			vars[key] = tVars.Get(key)
		} else {
			vars[fmt.Sprintf("%s%s", "dcim_", key)] = tVars.Get(key)
		}
	}
	vars["dcim_target_version"] = targetVersion
	vars["dcim_app_version"] = appVersion
	vars["dcim_component"] = component

	return &vars
}
