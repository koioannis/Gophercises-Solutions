package html_link_parser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func DFS(node *html.Node, links *[]Link) {
	href, text := parseNode(node)
	if len(href) > 0 {
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

		return href, strings.TrimSpace(string(text))
	}

	return "", ""
}

func parseTag(node *html.Node, text *[]byte) {
	if node.Type == html.TextNode {
		*text = append(*text, []byte(node.Data)...)
		*text = []byte(strings.TrimSpace(string(*text)))
		if !strings.HasSuffix(string(*text), " ") {
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
