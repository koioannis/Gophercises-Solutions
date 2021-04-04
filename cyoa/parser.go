package cyoa

import (
	"encoding/json"
	"io/ioutil"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Options
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p Parser) ParseJson(filepath string) (map[string]Arc, error) {
	var story map[string]Arc
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &story)
	return story, nil
}
