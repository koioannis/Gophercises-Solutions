package factory

import (
	"encoding/json"
)

type JSONParser struct{}

func newJSONParser() *JSONParser {
	return &JSONParser{}
}

func (jp JSONParser) Parse(data []byte) (map[string]string, error) {
	var rawData rawData
	err := json.Unmarshal(data, &rawData)
	if err != nil {
		return nil, err
	}

	return structToMap(rawData), nil
}
