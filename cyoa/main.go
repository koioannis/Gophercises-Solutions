package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var t *template.Template
var story map[string]Arc

func main() {
	story = parseJson()

	t, _ = template.ParseFiles("index.gohtml")

	myHandler := NewMyHandler()

	http.ListenAndServe(":3000", myHandler)
}

func NewMyHandler() *MyHanlder {
	return &MyHanlder{}
}

type MyHanlder struct{}

func (mh *MyHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t.Execute(w, story["intro"])
	}

	path := strings.Split(r.URL.Path, "/")[1]
	t.Execute(w, story[path])
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Options
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func parseJson() map[string]Arc {
	var story map[string]Arc
	data, err := ioutil.ReadFile(".\\data\\gopher.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(data, &story)

	return story
}
