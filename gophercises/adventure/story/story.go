package story

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Option of the story
type option struct {
	text string
	arc  string
}

//Story struct holds the story and its options
type Story struct {
	Title   string
	Story   []string
	Options []option
}

//Adventure holds all the stories
type Adventure map[string]Story

//GetStory returns the intro story
func (a Adventure) GetStory(name string) (Story, error) {
	intro, ok := a[name]
	if !ok {
		return intro, fmt.Errorf("Story %v, not found on the Adventure", name)
	}
	return intro, nil
}

//GetAdventure returns the stories for the adventure
func GetAdventure() Adventure {
	jsonAdventure := getJSON("gopher.json")

	adventure := make(Adventure)
	json.Unmarshal(jsonAdventure, &adventure)
	return adventure
}

func getJSON(file string) []byte {
	jf, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	encodedJSON, err := ioutil.ReadAll(jf)
	if err != nil {
		panic(err)
	}

	return encodedJSON
}
