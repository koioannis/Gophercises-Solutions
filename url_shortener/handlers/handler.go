package handlers

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		val, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, val, http.StatusSeeOther)
	}
}

type Config []struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	pathsToUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		val, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, val, http.StatusSeeOther)

	}, nil
}

// TODO: Refactor this, there must be a more performant way
func parseYAML(yml []byte) (map[string]string, error) {
	var config Config
	err := yaml.Unmarshal(yml, &config)
	if err != nil {
		return nil, err
	}

	urls := make(map[string]string)
	for index := range config {
		urls[config[index].Path] = config[index].Url
	}

	return urls, nil
}
