package main

import (
	"adventure/server"
	"adventure/story"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// 1. Get stories
	// 2. Load templates
	// 3. Fill templates with stories
	// 4. Handle stories path request
	// 5. Refactor to expose command interface
	// 6. Bonus points

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
