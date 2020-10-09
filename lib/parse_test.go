package markdown

import "testing"

// TestMarkdown tests the main Markdown function.
func TestMarkdown(t *testing.T) {
	heading := "## This is some text"
	html := Markdown(heading)

	t.Error("blah")

	println(html)
}
