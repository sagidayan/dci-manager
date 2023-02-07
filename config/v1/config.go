package v1

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rh-ecosystem-edge/dci-manager/config/common"
)

type Config struct {
	Version    string  `yaml:"version" json:"version"`
	Agents     []Agent `yaml:"agents" json:"agents"`
	suites     []common.Test
	prevConfig common.Configuration
}

var _ common.Configuration = &Config{}

func (c *Config) Plan() (common.ConfigurationDiff, error) {
	diff := c.diff()

	return diff, nil
}

func (c *Config) Apply() (common.ConfigurationDiff, error) {
	diff := c.diff()

	return diff, nil
}

func (c *Config) Sample(toJSON bool) ([]byte, error) {
	if !toJSON {
		return sampleYAML(), nil
	}
	return sampleJSON()
}

func (c *Config) Validate(data []byte, isYaml bool) error {
	// Schema check:
	err := common.ValidateSchema(configSchema, data, isYaml)
	if err != nil {
		return fmt.Errorf("Error: validation error. %s", err)
	}
	return nil
}

func (c *Config) IndexSuites() {
	c.suites = c.getSingleSuites()
}

func (c *Config) Tests() []common.Test {
	return c.suites
}

func (c *Config) MatrixString() string {
	resp := ""
	for _, agent := range c.Agents {
		t := table.NewWriter()
		t.SetTitle("%s  %s - Matrix", agent.Type, agent.ComponentName)
		t.AppendHeader(table.Row{"Target Version", "App Version", "Test"})
		suites := agent.SingleSuites()
		for _, suite := range suites {
			t.AppendRow(table.Row{suite.TargetVersion(), suite.AppVersion(), suite.Name()})
			t.AppendSeparator()
		}
		if len(resp) > 0 {
			resp = fmt.Sprintf("%s\n%s", resp, t.Render())
		} else {
			resp = t.Render()
		}
	}

	return resp
}

func (c *Config) ConfigVersion() string {
	return c.Version
}

func (c *Config) diff() common.ConfigurationDiff {
	var added = make([]common.Test, 0)
	var modified = make([]common.Test, 0)
	var removed = make([]common.Test, 0)
	if c.prevConfig == nil {
		added = c.suites
		return common.NewDiff(removed, added, modified)
	}

	removed = c.prevConfig.Tests()

	for _, newSuite := range c.suites {
		for i, oldSuite := range c.prevConfig.Tests() {
			if newSuite.Equals(oldSuite) {
				removed = common.RemoveSuiteFromSlice(removed, oldSuite)
				break
			}
			if newSuite.Modified(oldSuite) {
				removed = common.RemoveSuiteFromSlice(removed, oldSuite)
				modified = append(modified, newSuite)
				break
			}
			if len(c.prevConfig.Tests())-1 == i {
				added = append(added, newSuite)
			}
		}

	}

	return common.NewDiff(removed, added, modified)
}

func (c *Config) getSingleSuites() []common.Test {
	resp := make([]common.Test, 0)
	for _, agent := range c.Agents {
		resp = append(resp, agent.SingleSuites()...)
	}
	return resp
}

func (c *Config) SetPrevConfig(pc common.Configuration) {
	c.prevConfig = pc
}
