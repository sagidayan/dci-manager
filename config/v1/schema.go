package v1

import (
	_ "embed"
	"log"

	"sigs.k8s.io/yaml"
)

//go:embed schema.yaml
var fileSchemaYAML []byte

var configSchema string

func init() {
	b, err := yaml.YAMLToJSON(fileSchemaYAML)
	if err != nil {
		log.Fatal("Fatal: failed to intialize config scheme")
	}
	configSchema = string(b)
}
