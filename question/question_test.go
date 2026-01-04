package question

import (
	"bytes"
	"testing"
)

func TestAskQuestion(t *testing.T) {
	t.Run("test ask question", func(t *testing.T) {
		var (
			question = New([]string{"5+4", "9"})
			out      = bytes.NewBuffer(nil)
			in       = bytes.NewBufferString("9\n")
			got      = question.AskQuestion(out, in)
			want     = true
			gotOut   = out.String()
			wantOut  = "5+4: "
		)

		if wantOut != gotOut {
			t.Fatalf("want question '%s', got '%s", wantOut, gotOut)
		}
		if want != got {
			t.Fatal("answer is not correct")
		}
	})
}
