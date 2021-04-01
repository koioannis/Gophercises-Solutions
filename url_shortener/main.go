package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	mux := defaultMux()
	filename := initFlags()
	format := strings.Split(filepath.Ext(filename), ".")[1]

	handler, err := getHandler(format, filename, mux)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", handler)
}

func initFlags() string {
	inputFileMessage := `Provide the filename that contains the url paths. Note that it must be either json or yaml, and
	it must be under the data folder`

	var filename string
	flag.StringVar(&filename, "i", "map", inputFileMessage)
	flag.Parse()

	return filename
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi, try a link from your file")
	})
	return mux
}
