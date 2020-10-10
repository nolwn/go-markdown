package markdown

import "testing"

// TestMarkdown tests the main Markdown function.
func TestMarkdown(t *testing.T) {
	heading := "## This is some text\n"
	html := Markdown(heading)

	if html != "<h2>This is some text</h2>" {
		t.Error("## did not correctly generate an h2")
	}
}
