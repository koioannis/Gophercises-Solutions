package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := defaultMux()

	mapHandler, _ := getHandler("map", []byte{}, mux)

	yamlData, err := ioutil.ReadFile("data/urls.yaml")
	if err != nil {
		panic(err)
	}

	yamlHandler, err := getHandler("yaml", yamlData, mapHandler)
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
