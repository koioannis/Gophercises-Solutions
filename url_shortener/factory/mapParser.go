package factory

type MapParser struct{}


func newMapParser() *MapParser {
	return &MapParser{}
}

func (mp MapParser) Parse(data []byte) (map[string]string, error) {
	return map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}, nil
}