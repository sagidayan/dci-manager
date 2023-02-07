package common

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rh-ecosystem-edge/dci-manager/internal"
	"github.com/santhosh-tekuri/jsonschema/v3"
	"sigs.k8s.io/yaml"
)

func EqualSteps(s1 []TestStep, s2 []TestStep) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Order matters
	for i := range s1 {
		if s1[i].Render() != s2[i].Render() {
			return false
		}
	}
	return true
}

func EqualVariables(v1 TestVariables, v2 TestVariables) bool {
	if v1.Len() != v2.Len() {
		return false
	}

	for _, key := range v1.Keys() {
		if v1.Get(key) != v2.Get(key) {
			return false
		}
	}
	return true
}

func RemoveSuiteFromSlice(s []Test, suite Test) []Test {
	resp := []Test{}

	for _, candidet := range s {
		if !candidet.Equals(suite) {
			resp = append(resp, candidet)
		}
	}

	return resp
}

func ValidAgentTypes() []string {
	return []string{"openshift", "openshift-app", "rhel"}
}

func ValidateSchema(scheme string, data []byte, isYaml bool) error {
	var jsonObj []byte = data
	if isYaml {
		b, err := yaml.YAMLToJSON(data)
		if err != nil {
			return fmt.Errorf("Error: Invalid yaml file. %e", err)
		}
		jsonObj = b
	}

	var g interface{}
	err := json.Unmarshal(jsonObj, &g)
	if err != nil {
		return err
	}
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(internal.ConfigFile, strings.NewReader(scheme)); err != nil {
		return err
	}
	schema, err := compiler.Compile(internal.ConfigFile)
	if err != nil {
		return err
	}
	if err := schema.ValidateInterface(g); err != nil {
		if e, ok := err.(*jsonschema.ValidationError); ok {
			return fmt.Errorf("Location %s. Reason: %s", e.InstancePtr, e.Message)
		}
		return err
	}
	return nil

}
