package main

import (
	"koioannis/gopherices/url_shortener/factory"
	"net/http"
)

func getHandler(format string, data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parser, err := factory.GetParser(format)
	if err != nil {
		return nil, err
	}

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
