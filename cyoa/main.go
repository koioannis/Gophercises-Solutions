package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates *template.Template

func main() {
	filename := initFlags()

	story, err := NewParser().parseJson("data\\" + filename)
	if err != nil {
		fmt.Println("Please provide a valid filename (make sure your json is under the data folder)")
		os.Exit(0)
	}

	templates, _ = template.ParseFiles("views/index.gohtml", "views/notFound.gohtml")

	myHandler := NewMyHandler(story)

	fmt.Println()
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", myHandler)
}

func initFlags() string {
	inputFileMessage := "Provide the filename of the json file, including extension"

	var filename string
	flag.StringVar(&filename, "i", "", inputFileMessage)
	flag.Parse()

	return filename
}
