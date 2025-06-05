package embedutil

import (
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()
	got := List()
	want := []string{
		"templates/default.tmpl",
		"templates/system.tmpl",
		"themes/dark.json",
		"themes/light.json",
	}
	if len(got) != len(want) {
		t.Fatalf("want %d got %d", len(want), len(got))
	}
	for i, n := range want {
		if got[i] != n {
			t.Fatalf("want %s got %s", n, got[i])
		}
	}
}

func TestMustText(t *testing.T) {
	t.Parallel()
	txt := MustText("templates/default.tmpl")
	if !strings.Contains(txt, "{{.User}}") {
		t.Fatalf("missing variable in %q", txt)
	}
}

func TestMustTextPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	MustText("missing.txt")
}
