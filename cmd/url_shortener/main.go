package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/koioannis/gophercises-solutions/url_shortener"
)

func main() {
	mux := defaultMux()
	filename := initFlags()
	format := getFileExt(filename)

	handler, err := url_shortener.GetHandler(format, filename, mux)
	if err != nil {
		fmt.Println("WARNING: File not supported, defaults will be used")
	}

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", handler)
}

func getFileExt(filename string) string {
	fileExt := filepath.Ext(filename)

	if len(fileExt) != 0 {
		return strings.Split(fileExt, ".")[1]
	}

	fmt.Println("WARNING: File not supported, defaults will be used")
	return "map"
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
