package aiops

import "testing"

func TestTFIDFSummarizer(t *testing.T) {
    t.Parallel()
    s := NewTFIDFSummarizer()
    got := s.Summarize("error error fatal info info warning")
    if got == "" {
        t.Fatalf("empty summary")
    }
}
