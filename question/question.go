package question

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type Question struct {
	question,
	answer string
}

func New(r []string) Question {
	return Question{r[0], r[1]}
}

func (p Question) AskQuestion(out io.Writer, in io.Reader) bool {
	_, err := fmt.Fprintf(out, "%s: ", p.question)
	if err != nil {
		log.Fatalln("asking question:", err)
	}
	answer := readAnswer(in)
	return answer == p.answer
}

func readAnswer(in io.Reader) (answer string) {
	reader := bufio.NewReader(in)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("reading answer:", err)
	}
	return strings.TrimSpace(line)
}
