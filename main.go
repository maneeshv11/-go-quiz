package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {

	var fileName string
	var timeout int64
	var score int

	fmt.Printf("Welcome to quiz game\n")

	flag.StringVar(&fileName, "file", "problems.csv", "Path of the problem file.")
	flag.Int64Var(&timeout, "timeout", 30, "Maximum time for attempting each question")
	flag.Parse()

	problems, err := loadProblems(fileName)

	if err == nil {
	problemsLoop:
		for _, problem := range problems {
			fmt.Printf("%s : ", problem.question)

			ch := make(chan string)

			go expectedAnswer(ch)

			select {
			case answer := <-ch:
				if answer == problem.answer {
					score++
				}
			case <-time.After(time.Duration(timeout) * time.Second):
				fmt.Printf("Timeout.\n")
				break problemsLoop
			}

		}
	}

	fmt.Printf("Thanks for playing. You have correctly answered %d out of %d questions\n", score, len(problems))

}

func expectedAnswer(ch chan string) {
	var userAnswer string
	_, _ = fmt.Scanf("%s\n", &userAnswer)
	ch <- userAnswer
}

func loadProblems(problemFileName string) (problems []Problem, err error) {
	if problemFileName == "" {
		return nil, errors.New("file name is not valid")
	}

	csvFile, err := os.Open(problemFileName)

	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		problems = append(problems, Problem{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		})
	}

	return problems, nil

}
