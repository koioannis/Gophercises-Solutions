package factory

type JSONParser struct{}

func newJSONParser() *JSONParser {
	return &JSONParser{}
}

func (jp JSONParser) Parse(data []byte) (map[string]string, error) {

	return map[string]string{"test": "test"}, nil
}