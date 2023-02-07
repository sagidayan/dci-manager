package dci

import (
	"fmt"
	"log"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"github.com/rh-ecosystem-edge/dci-manager/dci/templates"
)

type settings struct {
	Topic     string
	Name      string
	JobFolder string
	Variables string
	filePath  string
}

func (s *settings) WriteToFile() error {
	err := templateParseToFile(s.filePath, templates.SettingsTemplate, s)
	if err != nil {
		return err
	}
	return nil
}
func newSettings(test common.Test) *settings {
	topic := "TOPIC"
	switch test.AgentName() {
	case DCIAgentOCP, DCIAgentOCPApp:
		topic = fmt.Sprintf("OCP-%s", test.TargetVersion())
	case DCIAgentRHEL:
		topic = fmt.Sprintf("RHEL-%s", test.TargetVersion())
	default:
		log.Fatalf("Error: invalid agent %s", test.AgentName())
	}
	s := &settings{
		Topic:     topic, //FIXME
		Name:      test.Name(),
		JobFolder: jobFolder(test),
		Variables: test.Variables().YAMLString(),
		filePath:  fmt.Sprintf("%s/settings.yml", jobFolder(test)),
	}
	return s
}
