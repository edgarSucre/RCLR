package problem

import (
	"encoding/csv"
	"os"
)

//Problem type encapsulates a question in the quiz
type Problem struct {
	q string
	a string
}

//ReadProblems returns a slice of questions
func ReadProblems(csvPath string) ([]Problem, error) {
	csv, err := readProblemCSV(csvPath)
	if err != nil {
		return nil, err
	}
	problems, err := parseCSV(csv)
	if err != nil {
		return nil, err
	}

	return problems, nil
}

func parseCSV(file *os.File) ([]Problem, error) {
	csvReader := csv.NewReader(file)
	csvContent, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	problems := make([]Problem, len(csvContent))
	for i, v := range csvContent {
		problems[i] = Problem{q: v[0], a: v[1]}
	}
	return problems, nil
}

func readProblemCSV(csvPath string) (*os.File, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	return f, nil
}
