package factory

// used to parse yaml and json
type config []struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type Parser interface {
	Parse([]byte) (map[string]string, error)
}

func structToMap(c config) map[string]string {
	urls := make(map[string]string)
	for index := range c {
		urls[c[index].Path] = c[index].Url
	}

	return urls
}