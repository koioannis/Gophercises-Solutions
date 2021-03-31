package factory

import "gopkg.in/yaml.v2"

type YAMLParser struct{}

func newYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

type config []struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// TODO: Refactor this, there must be a more performant way
func (yp YAMLParser) Parse(data []byte) (map[string]string, error) {
	var config config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	urls := make(map[string]string)
	for index := range config {
		urls[config[index].Path] = config[index].Url
	}

	return urls, nil
}
