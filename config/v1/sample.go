package v1

import (
	"bytes"
	_ "embed"
	"encoding/json"

	"sigs.k8s.io/yaml"
)

//go:embed sample.yaml
var fileSampleYAML []byte

func sampleYAML() []byte {
	return fileSampleYAML
}

func sampleJSON() ([]byte, error) {
	unindent, err := yaml.YAMLToJSONStrict(fileSampleYAML)
	if err != nil {
		return nil, err
	}

	b := bytes.Buffer{}
	err = json.Indent(&b, unindent, "", " ")
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
