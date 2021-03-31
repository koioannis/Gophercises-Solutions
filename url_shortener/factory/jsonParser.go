package factory

import (
	"encoding/json"
)

type JSONParser struct{}

func newJSONParser() *JSONParser {
	return &JSONParser{}
}

func (jp JSONParser) Parse(data []byte) (map[string]string, error) {
	var config config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return structToMap(config), nil
}
