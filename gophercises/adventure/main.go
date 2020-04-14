package main

import (
	"adventure/server"
	"adventure/story"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//TODO: apply functional options
	//TODO: remove panics from packages

	adventure := story.GetAdventure("./story/gopher.json")
	t := loadTemplate("index.html")

	handler := server.NewTaleHandler(t, adventure)

	log.Println(http.ListenAndServe(":8080", handler))
}

func loadTemplate(path string) *template.Template {
	t, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}
	return t
}
