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
)

type Problem struct {
	question string
	answer   string
}

func main() {

	var fileName string
	var score int

	fmt.Printf("Welcome to quiz game\n")

	flag.StringVar(&fileName, "file", "problems.csv", "Path of the problem file.")
	flag.Parse()

	problems, err := loadProblems(fileName)

	if err == nil {
		for _, problem := range problems {
			var userAnswer string
			fmt.Printf("%s : ", problem.question)

			_, _ = fmt.Scanf("%s", &userAnswer)

			if userAnswer == problem.answer {
				score++
			}

		}
	}

	fmt.Printf("Thanks for playing. You have correctly answered %d out of %d questions\n", score, len(problems))

}

func loadProblems(problemFileName string) (problems []Problem, err error) {
	if problemFileName == "" {
		return nil, errors.New("file name is not valid")
	}

	csvFile, err := os.Open(problemFileName)

	if err != nil {
		return nil, err
	}

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
			answer:   row[1],
		})
	}

	return problems, nil

}
