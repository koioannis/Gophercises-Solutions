package main

import (
	"strings"

	"golang.org/x/net/html"
)

func DFS(node *html.Node, links *[]Link) {
	href, text := parseNode(node)
	if len(href) > 0 {
		*links = append(*links, Link{Href: href, Text: text})
		node = node.LastChild
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

		return href, strings.TrimSpace(string(text))
	}

	return "", ""
}

func parseTag(node *html.Node, text *[]byte) {

	if node.Type == html.TextNode {
		*text = append(*text, []byte(node.Data)...)

		if !strings.HasSuffix(node.Data, " ") {
			*text = append(*text, ' ')
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
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
