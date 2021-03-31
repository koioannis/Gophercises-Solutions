package factory

import "errors"

func GetParser(parser string) (Parser, error) {
	if parser == "json" {
		return newJSONParser(), nil
	}

	if parser == "yaml" {
		return newYAMLParser(), nil
	}

	if parser == "map" {
		return newMapParser(), nil
	}

	return nil, errors.New("format not supported")
}