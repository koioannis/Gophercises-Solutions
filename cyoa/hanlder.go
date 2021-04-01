package main

import (
	"net/http"
	"strings"
)

type MyHanlder struct {
	story map[string]Arc
}

func NewMyHandler(story map[string]Arc) *MyHanlder {
	return &MyHanlder{
		story: story,
	}
}

func (mh *MyHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		templates.ExecuteTemplate(w, "index", mh.story["intro"])
		return
	}

	path := strings.Split(r.URL.Path, "/")[1]

	if arc, ok := mh.story[path]; ok {
		templates.ExecuteTemplate(w, "index", arc)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "notfound", nil)
}
