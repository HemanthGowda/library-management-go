package openapi

import (
	_ "embed"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

//go:embed openapi.yaml
var raw []byte

func JSON() ([]byte, error) {
	var document any
	if err := yaml.Unmarshal(raw, &document); err != nil {
		return nil, err
	}
	return json.Marshal(document)
}
