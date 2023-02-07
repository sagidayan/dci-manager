package dci

import (
	"fmt"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"github.com/rh-ecosystem-edge/dci-manager/dci/templates"
)

type writeAbleTemplate interface {
	WriteToFile(filePath string) error
}

type ocpRunScript struct {
	Name           string
	JobFolder      string
	KubeconfigPath string //FIXME
}

type rhelRunScript struct {
	Name      string
	JobFolder string
}

type ocpAppRunScript struct {
	Name           string
	JobFolder      string
	KubeconfigPath string //FIXME
}

var _ writeAbleTemplate = &rhelRunScript{}
var _ writeAbleTemplate = &ocpRunScript{}
var _ writeAbleTemplate = &ocpAppRunScript{}

func newRunScript(test common.Test) (writeAbleTemplate, error) {
	var r writeAbleTemplate
	kubeconfig := test.Variables().Get("dcim_kubeconfig_path")
	switch test.AgentName() {
	case DCIAgentOCPApp:
		r = &ocpAppRunScript{
			Name:           test.Name(),
			JobFolder:      jobFolder(test),
			KubeconfigPath: kubeconfig, //FIXME
		}
	case DCIAgentOCP:
		r = &ocpRunScript{
			Name:           test.Name(),
			JobFolder:      jobFolder(test),
			KubeconfigPath: kubeconfig, //FIXME
		}
	case DCIAgentRHEL:
		r = &rhelRunScript{
			Name:      test.Name(),
			JobFolder: jobFolder(test),
		}
	default:
		return nil, fmt.Errorf("Invalid agent type. %s/%s/%s are allowed, got %s", DCIAgentOCP, DCIAgentRHEL, DCIAgentOCPApp, test.AgentName())

	}

	return r, nil
}

func (s *ocpRunScript) WriteToFile(filePath string) error {
	err := templateParseToFile(filePath, templates.RunScriptTemplateOpenShift, s)
	if err != nil {
		return err
	}
	return nil
}
func (s *rhelRunScript) WriteToFile(filePath string) error {
	err := templateParseToFile(filePath, templates.RunScriptTemplateRHEL, s)
	if err != nil {
		return err
	}
	return nil
}
func (s *ocpAppRunScript) WriteToFile(filePath string) error {
	err := templateParseToFile(filePath, templates.RunScriptTemplateOpenShiftApp, s)
	if err != nil {
		return err
	}
	return nil
}
