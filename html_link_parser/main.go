package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func main() {
	file, err := os.Open("data/ex1.html")
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}

	var links []Link
	DFS(doc, &links)

	jsonFile, _ := json.MarshalIndent(links, "", "")

	_ = ioutil.WriteFile("out.json", jsonFile, 0644)

}

func DFS(node *html.Node, links *[]Link) {
	href, text := parseNode(node)
	if len(href) > 0 {
		fmt.Println(href, text)
		*links = append(*links, Link{Href: href, Text: text})
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		DFS(c, links)
	}

}

func parseNode(node *html.Node) (string, string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		href := strings.TrimSpace(parseAttr(node.Attr))

		if node.FirstChild == nil {
			return href, ""
		}

		var text []byte
		parseTag(node, &text)

		return href, string(text)
	}

	return "", ""
}

func parseTag(node *html.Node, text *[]byte) {

	if node.Type == html.TextNode {
		*text = append(*text, []byte(node.Data)...)
		*text = append(*text, ' ')
	}

	for c := node.FirstChild; c != nil; c =	c.NextSibling {
		parseTag(c, text)
	}
}

func parseAttr(attr []html.Attribute) string {
	var href string

	for _, val := range attr {
		if val.Key == "href" {
			href = val.Val
		}
	}

	return href
}
