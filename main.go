package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	problemsFile = "problems.csv"
	welcomeMsg   = "Welcome to the Quiz! Press Enter to start."
	resultMsg    = "You got %d out of %d questions correct."
)

func main() {
	timeLimitPtr := flag.Int("timeLimit", 2, "time limit for questionnaire in seconds")

	flag.Parse()

	records := readCSV(problemsFile)
	totalQuestions := len(records)

	fmt.Println(welcomeMsg)
	bufio.NewReader(os.Stdin).ReadString('\n') // Wait for user to press Enter

	timeLimit := time.Duration(*timeLimitPtr) * time.Second
	totalCorrect := questionnaire(records, timeLimit)

	fmt.Printf(resultMsg, totalCorrect, totalQuestions)
}

func questionnaire(records [][]string, timeLimit time.Duration) int {
	totalCorrect := 0
	timer := time.After(timeLimit)

	for _, record := range records {
		select {
		case <-timer:
			fmt.Println("\nTime's up!")
			return totalCorrect
		default:
			question, answer := record[0], strings.TrimSpace(record[1])
			if askQuestion(question) == answer {
				totalCorrect++
			}
		}
	}

	return totalCorrect
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't open '%s': %s", filename, err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse the CSV file: %s", err.Error())
	}

	return records
}

func askQuestion(question string) string {
	fmt.Printf("%s? ", question)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %s", err.Error())
	}
	return strings.TrimSpace(input)
}
