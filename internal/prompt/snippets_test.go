package prompt

import "testing"

func TestSnippetsNotEmpty(t *testing.T) {
	if Bash == "" || Zsh == "" || Fish == "" || PowerShell == "" {
		t.Fatalf("snippets empty")
	}
}
