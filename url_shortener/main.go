package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"koioannis/gopherices/url_shortener/handlers"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	yamlData, err := ioutil.ReadFile("data/urls.yaml")
	if err != nil {
		panic(err)
	}
	_, err = handlers.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := handlers.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
