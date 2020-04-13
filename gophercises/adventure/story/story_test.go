package story_test

import (
	"adventure/story"
	"testing"
)

func TestGetStories(t *testing.T) {
	adventure := story.GetAdventure()
	_, err := adventure.GetStory("intro")
	if err != nil {
		t.Error("Story Intro expected")
	}
}
