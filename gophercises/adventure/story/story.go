package story

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Option of the story
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//Story struct holds the story and its options
type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//Adventure holds all the stories
type Adventure map[string]Story

//GetStory returns the intro story
func (a *Adventure) GetStory(name string) (Story, error) {
	intro, ok := (*a)[name]
	if !ok {
		return intro, fmt.Errorf("Story '%v', not found on the Adventure", name)
	}
	return intro, nil
}

//GetAdventure returns the stories for the adventure
func GetAdventure(filePath string) (*Adventure, error) {
	jsonAdventure, err := getJSON(filePath)
	if err != nil {
		return nil, err
	}

	adventure := new(Adventure)
	json.Unmarshal(jsonAdventure, &adventure)
	return adventure, nil
}

func getJSON(file string) ([]byte, error) {
	jf, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	encodedJSON, err := ioutil.ReadAll(jf)
	if err != nil {
		return nil, err
	}

	return encodedJSON, nil
}
