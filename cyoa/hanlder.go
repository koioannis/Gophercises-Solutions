package cyoa

import (
	"net/http"
	"strings"
	"html/template"
)

type MyHanlder struct {
	story     map[string]Arc
	templates *template.Template
}

func NewMyHandler(story map[string]Arc, templates *template.Template) *MyHanlder {
	return &MyHanlder{
		story:     story,
		templates: templates,
	}
}

func (mh *MyHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		mh.templates.ExecuteTemplate(w, "index", mh.story["intro"])
		return
	}

	path := strings.Split(r.URL.Path, "/")[1]

	if arc, ok := mh.story[path]; ok {
		mh.templates.ExecuteTemplate(w, "index", arc)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	mh.templates.ExecuteTemplate(w, "notfound", nil)
}
