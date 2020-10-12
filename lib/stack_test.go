package markdown

import (
	"fmt"
	"testing"
)

var e1 = element{}
var e2 = element{kind: "tag", isClosing: true, tag: "h1"}
var e3 = element{kind: "text", text: "This is some text"}

func TestPush(t *testing.T) {
	s := elementStack{}

	s.push(e1)
	s.push(e2)
	s.push(e3)
}

func TestPop(t *testing.T) {
	s := elementStack{}

	s.push(e1)
	ok, pop := s.pop()

	if !ok {
		t.Error("The stack should have popped, but did not.")
	}

	if pop.kind != "" {
		t.Error("p1 kind should be an empty string")
	}

	ok, pop = s.pop()

	fmt.Printf("%v", pop)
	if ok {
		t.Error("The stack should have been empty, but it popped.")
	}
}
