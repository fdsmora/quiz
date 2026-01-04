package quiz

import (
	"fmt"
	"io"

	"github.com/fdsmora/gophercises/quiz/question"
)

type Quiz struct {
	questions      []question.Question
	correctAnswers int
}

func New(questions []question.Question) Quiz {
	return Quiz{
		questions: questions,
	}
}

func (q *Quiz) Run(out io.Writer, in io.Reader) {
	fmt.Fprintln(out, "Welcome! Press any key to continue...")
	x := ""
	fmt.Fscanln(in, &x)

	for _, qst := range q.questions {
		if qst.AskQuestion(out, in) {
			q.correctAnswers++
		}
	}
}

func (q Quiz) PrintResult(out io.Writer) {
	fmt.Fprintf(out,
		"You got '%d' answers out of '%d'\n",
		q.correctAnswers,
		len(q.questions))
}
