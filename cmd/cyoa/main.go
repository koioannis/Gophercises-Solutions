package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"koioannis/gophercises/cyoa"
)

func main() {
	filename := initFlags()

	story, err := cyoa.NewParser().ParseJson("cyoa\\data\\" + filename)
	if err != nil {
		fmt.Println("Please provide a valid filename (make sure your json is under the data folder)")
		os.Exit(0)
	}

	templates, _ := template.ParseFiles("cyoa\\views\\index.gohtml", "cyoa\\views\\notFound.gohtml")

	myHandler := cyoa.NewMyHandler(story, templates)

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
