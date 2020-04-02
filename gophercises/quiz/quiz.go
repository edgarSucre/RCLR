package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func (p *problem) checkAnswer(a string, right, wrong *int) {
	if p.answer == a {
		*right++
		return
	}
	*wrong++
}

func main() {
	csvPath, limit := parseFlags()
	csvFile := openCSV(csvPath)
	problems := readQuestions(csvFile)
	right, wrong := askQuestions(&problems, limit, bufio.NewReader(os.Stdin))
	fmt.Println("Got right:", right)
	fmt.Println("Got worng:", wrong)
}

func askQuestions(problems *[]problem, limit int, reader *bufio.Reader) (right, wrong int) {
	timer := time.NewTimer(time.Duration(limit) * time.Second).C
	input := make(chan string)

	for _, p := range *problems {
		fmt.Print(p.question + ": ")
		go getInput(input, reader)
		select {
		case <-timer:
			fmt.Print("\nTime's UP!!\n")
			return
		case ans := <-input:
			p.checkAnswer(ans, &right, &wrong)
		}
	}
	return
}

func readQuestions(f *os.File) []problem {
	reader := csv.NewReader(f)
	var problems []problem
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		problems = append(problems, problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		})
	}
	return problems
}

func getInput(input chan string, prompt *bufio.Reader) {
	guess, err := prompt.ReadString('\n')

	if err != nil {
		fmt.Println("failed to read answer")
		input <- ""
	}
	guess = strings.TrimSpace(guess)
	input <- guess
}

func openCSV(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic("Could not open file in path: " + path)
	}
	return file
}

func parseFlags() (string, int) {
	path := flag.String("csv", "./problems.csv",
		`a csv file path, default: "problems.csv", format: "question, answer" `)
	l := flag.Int("limit", 60, "time limit in secods to solve the quiz default 60")
	flag.Parse()
	return *path, *l
}
