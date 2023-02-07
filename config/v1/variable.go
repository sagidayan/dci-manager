package v1

import (
	"github.com/rh-ecosystem-edge/dci-manager/config/common"
	"gopkg.in/yaml.v2"
)

type variables map[string]string

var _ common.TestVariables = &variables{}

const (
	keyPrefix = "dcim_"
)

func (v *variables) Get(key string) string {
	m := *v
	val, ok := m[key]
	if !ok {
		return ""
	}
	return val
}

func (v *variables) Len() int {
	m := *v
	return len(m)
}

func (v *variables) Keys() []string {
	m := *v
	resp := []string{}
	for key := range m {
		resp = append(resp, key)
	}
	return resp
}

func (v *variables) Values() []string {
	m := *v
	resp := []string{}
	for _, val := range m {
		resp = append(resp, val)
	}
	return resp
}

func (v *variables) YAMLString() string {
	b, _ := yaml.Marshal(v)
	return string(b)
}
