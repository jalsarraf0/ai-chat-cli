package aiops

import "testing"

func TestRegexDetector(t *testing.T) {
    t.Parallel()
    det, err := NewRegexDetector([]string{"ERROR", "FATAL"})
    if err != nil {
        t.Fatalf("new: %v", err)
    }
    tests := []struct{
        line string
        ok   bool
    }{
        {"info: ok", false},
        {"ERROR: bad", true},
        {"FATAL: crash", true},
    }
    for _, tt := range tests {
        tt := tt
        t.Run(tt.line, func(t *testing.T) {
            t.Parallel()
            _, got := det.Detect(nil, tt.line)
            if got != tt.ok {
                t.Fatalf("Detect=%v want %v", got, tt.ok)
            }
        })
    }
}
