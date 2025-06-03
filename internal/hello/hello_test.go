package hello

import "testing"

func TestGreet(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"default", "", "hello, world"},
		{"named", "Codex", "hello, Codex"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := Greet(c.input); got != c.want {
				t.Fatalf("want %q got %q", c.want, got)
			}
		})
	}
}
