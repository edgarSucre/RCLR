package main

import (
	"bufio"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	path, limit := parseFlags()
	if path != "./problems.csv" || limit != 60 {
		t.Error("asdasd")
	}
}

func TestReadQuestions(t *testing.T) {
	f := openCSV("./problems.csv")
	p := readQuestions(f)

	if p[0].question != "5+5" {
		t.Error("Failed reading problems")
	}
}

func TestAskQuestions(t *testing.T) {
	f := openCSV("./problems.csv")
	p := readQuestions(f)

	f, err := os.Open("testAnswers.txt")
	if err != nil {
		t.Error("Could not open the aswers file")
	}
	reader := bufio.NewReader(f)
	right, wrong := askQuestions(&p, 200, reader)
	if right != 12 || wrong != 0 {
		t.Error("Right and Wrong does not match")
	}
}

func TestOpenCSV(t *testing.T) {
	f := openCSV("./problems.csv")
	if f.Name() == "problems.csv" {
		t.Error("File name does not match")
	}
	defer stopCrash(t)
	_ = openCSV("no-file.csv")
}

func stopCrash(t *testing.T) {
	errMsg := recover()
	err, ok := errMsg.(string)
	if ok {
		if err != "Could not open file in path: no-file.csv" {
			t.Error("Error message does not match")
		}
	} else {
		t.Error("TestOpenCSV did't panic with a string msg")
	}
}
