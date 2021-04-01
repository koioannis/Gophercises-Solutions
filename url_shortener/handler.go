package main

import (
	"io/ioutil"
	"koioannis/gopherices/url_shortener/factory"
	"net/http"
)

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getHandler(format string, filename string, fallback http.Handler) (http.HandlerFunc, error) {
	parser, err := factory.GetParser(format)
	if err != nil || format == "map" {
		parser, _ = factory.GetParser("map")
		return buildHandler(parser, nil, fallback)
	}

	data, err := readFile(filename)
	if err != nil {
		parser, _ = factory.GetParser("map")
	}

	return buildHandler(parser, data, fallback)
}

func buildHandler(parser factory.Parser, data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := parser.Parse(data)
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
