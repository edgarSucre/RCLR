package main

import (
	"adventure/server"
	"adventure/story"
	"log"
	"net/http"
)

func main() {
	adventure, err := story.GetAdventure("./story/gopher.json")
	if err != nil {
		panic(err)
	}

	handler := server.NewTaleHandler(adventure)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
