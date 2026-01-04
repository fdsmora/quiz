package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fdsmora/gophercises/quiz/question"
	"github.com/fdsmora/gophercises/quiz/quiz"
)

const (
	problemsFile = "problems.csv"
)

func main() {
	timeLimitPtr := flag.Int("timeLimit", 2, "time limit for questionnaire in seconds")
	shuffle := false
	var shufflePtr *bool = &shuffle
	flag.BoolVar(shufflePtr, "s", false, "shorthand for --shuffle")

	flag.Parse()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	records := readCSV(problemsFile)

	totalQuestions := len(records)
	if *shufflePtr {
		// Shuffle the slice
		r.Shuffle(totalQuestions, func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	quiz := quiz.New(createQuestions(records))
	go quiz.Run(os.Stdout, os.Stdin)

	// Wait for the quiz to finish
	time.Sleep(time.Duration(*timeLimitPtr) * time.Second)
	quiz.PrintResult(os.Stdout)
}

func createQuestions(records [][]string) []question.Question {
	questions := make([]question.Question, 0, len(records))
	for _, r := range records {
		questions = append(questions, question.New(r))
	}
	return questions
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open '%s': %s", filename, err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("error reading CSV: %s", err.Error())
		}
		records = append(records, record)
	}

	return records
}
