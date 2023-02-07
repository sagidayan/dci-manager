package v1

import (
	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"sigs.k8s.io/yaml"
)

//type Step struct {
//	Name        string `yaml:"name" json:"name"`
//	AnsibleTask string `yaml:"ansible_task" json:"ansible_task"`
//}

type Step map[string]any

var _ common.TestStep = &Step{}

func (s *Step) Render() string {
	//TODO: Implement!
	b, _ := yaml.Marshal(s)
	return string(b)
}
