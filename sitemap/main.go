package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url, err := initFlags()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := getResponseFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(node)
}

func getResponseFromURL(url string) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("can't fetch requested URL (status code %v)", resp.StatusCode)
	}

	return resp, nil
}

func initFlags() (string, error) {
	URLDescription := "The url you want to build a sitemap"

	var inURL string
	flag.StringVar(&inURL, "url", "", URLDescription)
	flag.Parse()

	if len(inURL) == 0 {
		fmt.Println("You must provide a url")
		return "", errors.New("Empty URL tag")
	}

	return inURL, nil
}
