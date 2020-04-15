package story_test

import (
	"adventure/story"
	"testing"
)

func TestGetAdventure(t *testing.T) {
	t.Run("Invalid file path", func(t *testing.T) {
		_, err := story.GetAdventure("")
		if err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("Valid file path", func(t *testing.T) {
		adventure, _ := story.GetAdventure("./gopher.json")
		if adventure == nil {
			t.Error("Adventure expected")
		}
	})

}

func TestGetSotry(t *testing.T) {
	adventure, err := story.GetAdventure("./gopher.json")
	if err != nil {
		t.Error(err)
	}

	t.Run("Invalid Story", func(t *testing.T) {
		_, err := adventure.GetStory("")
		if err == nil {
			t.Error("Error expected")
		}
	})

	t.Run("Valid Story", func(t *testing.T) {
		intro, err := adventure.GetStory("intro")
		if err != nil {
			t.Error("Story Intro expected")
		}
		if len(intro.Options) == 0 || len(intro.Story) == 0 || len(intro.Title) == 0 {
			t.Error("Bad formed story")
		}
	})
}
