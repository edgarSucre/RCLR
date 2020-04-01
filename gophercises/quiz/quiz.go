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

var csvPath string
var limit int

func init() {
	flag.StringVar(&csvPath, "csv", "./problems.csv",
		`a csv file path, default: "problems.csv", format: "question, answer" `)
	flag.IntVar(&limit, "limit", 60,
		"time limit in secods to solve the quiz default 60")
	flag.Parse()
}

func main() {
	csvFile := openCSV(csvPath)
	csvReader := csv.NewReader(csvFile)
	right, wrong := askQuestions(csvReader)
	fmt.Println("Got right:", right)
	fmt.Println("Got worng:", wrong)
}

func askQuestions(reader *csv.Reader) (right, wrong int) {
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	input := make(chan string)

	for {
		question, err := reader.Read()
		if err != nil {
			break
		}
		for {
			go getInput(question[0], input)
			select {
			case <-timer.C:
				fmt.Println()
				return
			case in := <-input:
				if in == question[1] {
					right++
				} else {
					wrong++
				}
			}
			break
		}
	}
	return
}

func getInput(question string, input chan string) {
	prompt := bufio.NewReader(os.Stdin)

	fmt.Print(question + ": ")
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
