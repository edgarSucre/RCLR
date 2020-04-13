package main

import (
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

	adventure := story.GetAdventure()

	intro, err := adventure.GetStory("intro")
	if err != nil {
		panic(err)
	}

	serve(intro)
}

func serve(intro story.Story) {
	handler := func(response http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		t.Execute(response, intro)
	}

	http.HandleFunc("/", handler)
	log.Println("Listening on port 8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
