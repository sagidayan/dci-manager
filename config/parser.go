package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"

	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"github.com/rh-ecosystem-edge/dci-manager/config/v1"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
)

func Parse(path string) (common.Configuration, error) {
	c, err := parseConfig(path)
	if err != nil {
		return nil, err
	}
	prevConf, err := parseConfig(internal.StateFilePath)
	if err == nil && prevConf != nil {
		prevConf.IndexSuites()
		c.SetPrevConfig(prevConf)
	}
	c.IndexSuites()
	return c, nil
}

func parseConfig(path string) (common.Configuration, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file '%s'", internal.ConfigFile)
	}
	basic := &bareConfig{}
	ext := filepath.Ext(path)
	isYaml := true

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(dat, basic)
	case ".json":
		isYaml = false
		err = json.Unmarshal(dat, basic)
	default:
		return nil, fmt.Errorf("%s invalid file type. Expected json/yaml", internal.ConfigFile)

	}
	if err != nil {
		return nil, fmt.Errorf("Error: failed to parse %s. %s", internal.ConfigFile, err)
	}

	var c common.Configuration
	switch basic.Version {
	case "1":
		c = &v1.Config{}
		err := c.Validate(dat, isYaml)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(dat, c)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Error: parsing error. unknown config version '%s'. provided config: %s", basic.Version, internal.ConfigFile)
	}

	return c, nil
}

type bareConfig struct {
	Version string `yaml:"version" json:"version"`
}
