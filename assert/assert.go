package assert

import (
	"bytes"
	"log"
	"testing"

	"github.com/alecthomas/chroma/quick"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func Highlight(s string) string {
	buf := bytes.NewBufferString("")
	if err := quick.Highlight(buf, s, "bash", "terminal256", "monokai"); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func Equal(t *testing.T, expected string, actual string) {
	if expected == actual {
		t.Log(Highlight(actual))
	} else {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, actual, false)
		t.Errorf("\nexpected: %v\nactual  : %v", expected, dmp.DiffPrettyText(diffs))
	}
}
