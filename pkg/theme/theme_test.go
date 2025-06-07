package theme

import (
	"os"
	"testing"
)

func TestLoadDefault(t *testing.T) {
	t.Parallel()
	os.Setenv("COLORTERM", "light")
	p := Load("")
	if p.Background != "#ffffff" {
		t.Fatalf("want light palette")
	}
}
