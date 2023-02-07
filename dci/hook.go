package dci

import (
	"fmt"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"github.com/rh-ecosystem-edge/dci-manager/dci/templates"
	"gopkg.in/yaml.v2"
)

type HookType string

const (
	HookTypePreRun   HookType = "pre-run"
	HookTypePostRun  HookType = "post-run"
	HookTypeTests    HookType = "tests"
	HookTypeInstall  HookType = "install"
	HookTypeTeardown HookType = "teardown"
)

type hook struct {
	hookType HookType
}

func (h *hook) WriteHook(folder string, steps []common.TestStep) error {
	s, _ := yaml.Marshal(steps)
	if len(steps) == 0 {
		s = []byte{}
	}
	fileName := fmt.Sprintf("%s.yml", h.hookType)
	fileLocation := fmt.Sprintf("%s/%s", folder, fileName)
	err := templateParseToFile(fileLocation, templates.HookTemplate, string(s))
	if err != nil {
		return err
	}
	return nil
}
