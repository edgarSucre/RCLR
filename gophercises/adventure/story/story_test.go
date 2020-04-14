package story_test

import (
	"adventure/story"
	"testing"
)

func TestGetStories(t *testing.T) {
	adventure := story.GetAdventure("./gopher.json")
	intro, err := adventure.GetStory("intro")
	if err != nil {
		t.Error("Story Intro expected")
	}
	if len(intro.Options) == 0 || len(intro.Story) == 0 || len(intro.Title) == 0 {
		t.Error("Bad formed story")
	}
}
