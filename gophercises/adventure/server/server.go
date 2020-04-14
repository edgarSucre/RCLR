package server

import (
	"adventure/story"
	"html/template"
	"net/http"
	"strings"
)

//Tale is the handler for the story
type Tale struct {
	template  *template.Template
	adventure story.Adventure
}

func (t *Tale) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := strings.TrimPrefix(req.URL.Path, "/")
	if key == "" {
		key = "intro"
	}
	if key != "favicon.ico" {
		story, err := t.adventure.GetStory(key)
		if err != nil {
			panic(err)
		}
		t.template.Execute(resp, story)
	}
	return
}

//NewTaleHandler returns a pointer to the Tale server
func NewTaleHandler(temp *template.Template, ad story.Adventure) *Tale {
	return &Tale{template: temp, adventure: ad}
}
