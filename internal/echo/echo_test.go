package echo

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		n     int
		want  string
	}{
		{"single", "a", 1, "a"},
		{"double", "b", 2, "bb"},
		{"triple", "c", 3, "ccc"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Repeat(tt.input, tt.n); got != tt.want {
				t.Fatalf("expected %q got %q", tt.want, got)
			}
		})
	}
}
