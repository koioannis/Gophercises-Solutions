package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"koioannis/gophercises/html_link_parser"

	"golang.org/x/net/html"
)

func main() {
	inputFilename, outFilename := initFlags()

	file := openHtmlFile(inputFilename)

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	var links []html_link_parser.Link
	html_link_parser.DFS(doc, &links)

	writeJson(outFilename, links)

}

func openHtmlFile(filename string) io.Reader {
	file, err := os.Open("html_link_parser/data/" + filename)
	if err != nil {
		panic(err)
	}

	return file
}

func writeJson(filename string, links []html_link_parser.Link) {
	jsonFile, err := json.MarshalIndent(links, "", "")
	if err != nil {
		panic(err)
	}

	path := "html_link_parser/" + filename 
	err = ioutil.WriteFile(path, jsonFile, 0644)
	if err != nil {
		panic(err)
	}
}

func initFlags() (string, string) {
	inputFileDescription := "The html file you want to parse, including extension"
	outFileDescription := "The json file you want to save the parsed tags, including extensions"

	var filename, outFilename string
	flag.StringVar(&filename, "i", "", inputFileDescription)
	flag.StringVar(&outFilename, "o", "", outFileDescription)

	flag.Parse()

	if filepath.Ext(filename) != ".html" {
		fmt.Println("Please provide a valid input file (extension included)")
		os.Exit(0)
	}

	if filepath.Ext(outFilename) != ".json" {
		fmt.Println("Please provide a valid out file (extension included)")
		os.Exit(0)
	}

	return filename, outFilename
}
