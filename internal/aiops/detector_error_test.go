package aiops

import "testing"

func TestNewRegexDetectorError(t *testing.T) {
	if _, err := NewRegexDetector([]string{"["}); err == nil {
		t.Fatalf("expected error")
	}
}
