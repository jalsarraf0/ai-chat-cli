package shell

import "testing"

func TestDetect(t *testing.T) {
	t.Parallel()
	if got := Detect(); got != "bash" {
		t.Fatalf("Detect()=%q", got)
	}
}
