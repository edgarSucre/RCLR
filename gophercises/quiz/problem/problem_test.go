package problem_test

import (
	"errors"
	"os"
	"quiz/problem"
	"testing"
)

func TestReadQuestions(t *testing.T) {
	t.Run("Read from valid file", func(t *testing.T) {
		problems, err := problem.ReadProblems("../problems.csv")
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(problems) != 12 {
			t.Fatal("Problems len should be 12")
		}
	})

	t.Run("Read from invalid file", func(t *testing.T) {
		_, err := problem.ReadProblems("../problemsss.csv")
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal("error missmatch")
		}
	})
}
