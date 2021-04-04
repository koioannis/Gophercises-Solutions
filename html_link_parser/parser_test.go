package html_link_parser

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func createNode(htmlString string) (node *html.Node) {
	node, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}

	return
}

func areEqual(expected []Link, actual []Link) bool {

	for i := 0; i < len(expected); i++ {
		if !reflect.DeepEqual(expected[i], actual[i]) {
			return false
		}
	}
	return true
}

func Test_DFS_WhenGivenHtmlWithoutAnchorTags_ShouldReturnEmptyLink(t *testing.T) {
	expected := []Link{}

	htmlString := `
	<p>Hi, im html</p>
	<h1>Im an h1 tag</h1>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}

func Test_DFS_WhenGivenOneAnchorTags_ShouldReturnTheTextInsideIt(t *testing.T) {
	expected := []Link{
		{
			Href: "/golang",
			Text: "Go is awesome!",
		},
	}

	htmlString := `
	<a href="/golang">Go is awesome!</a>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}

func Test_DFS_WhenGivenMultipleAnchorTags_ShouldReturnTheTextInsideThem(t *testing.T) {
	expected := []Link{
		{
			Href: "/golang",
			Text: "Go is awesome!",
		},
		{
			Href: "https://github.com/koioannis",
			Text: "That's my github",
		},
		{
			Href: "https://go.dev/",
			Text: "Great docs",
		},
	}

	htmlString := `
	<a href="/golang">Go is awesome!</a>
	<a href="https://github.com/koioannis">That's my github</a>
	<a href="https://go.dev/">Great docs</a>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}

func Test_DFS_WhenGivenNestedAnchorTags_ShouldReturnTheNestedAnchorTags(t *testing.T) {
	expected := []Link{
		{
			Href: "/golang",
			Text: "Go is awesome!",
		},
		{
			Href: "https://github.com/koioannis",
			Text: "That's my github",
		},
		{
			Href: "https://go.dev/",
			Text: "Great docs",
		},
	}

	htmlString := `
	<a href="/golang">Go is awesome! <a href="https://github.com/koioannis"> That's my github</a> </a>
	<a href="https://go.dev/">Great docs</a>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}

func Test_DFS_WhenGivenAnchorTagsWithNestedTags_ShouldReturnTheTextOfTheNestedTags(t *testing.T) {
	expected := []Link{
		{
			Href: "/dog",
			Text: "Something in a span Text not in a span Bold text!",
		},
	}

	htmlString := `
	<a href="/dog">
	<span>Something in a span</span>Text not in a span
	<b>Bold text!</b></a>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}

func Test_DFS_WhenGivenAnchorTagsWithNestedAnchorTags_ShouldNotReturnTheTextOfNested(t *testing.T) {
	expected := []Link{
		{
			Href: "/parent",
			Text: "This is text",
		},
		{
			Href: "/nested",
			Text: "This is nested text",
		},
	}

	htmlString := `
	<a href="/parent">This is text<a href="/nested">This is nested text</a></a>`
	node := createNode(htmlString)

	var links []Link
	DFS(node, &links)
	if !areEqual(expected, links) {
		t.Errorf("Expected [%s] to equal [%s]", links, expected)
	}
}
