package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
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

	fmt.Println(welcomeMsg)
	bufio.NewReader(os.Stdin).ReadString('\n') // Wait for user to press Enter

	timeLimit := time.Duration(*timeLimitPtr) * time.Second
	totalCorrect := questionnaire(records, timeLimit)

	fmt.Printf(resultMsg, totalCorrect, totalQuestions)
}

func askQuestion(question string, aCh chan string, doneCh chan bool) {
	fmt.Printf("%s? ", question)

	inputCh := make(chan string)
	go func() {
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %s", err.Error())
		}
		inputCh <- input
	}()

	select {
	case input := <-inputCh:
		aCh <- input
	case <-doneCh:
		return
	}
}

func questionnaire(records [][]string, timeLimit time.Duration) int {
	totalCorrect := 0

	timer := time.After(timeLimit) // Single timer for the entire quiz
	aCh := make(chan string)
	doneCh := make(chan bool)

	for _, record := range records {
		question, answer := strings.TrimSpace(record[0]), strings.TrimSpace(record[1])
		go func() {
			askQuestion(question, aCh, doneCh)
		}()

		select {
		case <-timer:
			close(doneCh) // Signal all goroutines to stop
			fmt.Println("\nTime's up! totalCorrect:", totalCorrect)
			return totalCorrect
		case ans := <-aCh:
			if strings.TrimSpace(ans) == answer {
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
