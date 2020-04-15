package server

import (
	"adventure/story"
	"html/template"
	"net/http"
	"strings"
)

var defaultTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose your adventure</title>
    <style>
        body {
            margin: 0 10%;
        }
    </style>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Story}}
        <p>{{.}}</p>
    {{end}}

    <h3>Choose your fate</h3>
    <ul>
        {{range .Options}}
            <li><a href="/{{.Arc}}">{{.Text}}</a></li>        
        {{end}}
    </ul>    
</body>
</html>`

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

type tale struct {
	template  *template.Template
	adventure *story.Adventure
}

//TaleOption dinamically sets options for the TaleHandler
type TaleOption func(t *tale)

//WithTemplate sets a template for the tale
func WithTemplate(temp *template.Template) TaleOption {
	return func(t *tale) {
		t.template = temp
	}
}

//NewTaleHandler returns a story handler
func NewTaleHandler(ad *story.Adventure, opts ...TaleOption) http.Handler {
	handler := tale{adventure: ad, template: tpl}
	for _, opt := range opts {
		opt(&handler)
	}
	return handler
}

func (t tale) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
