//go:build unit

package openai

import (
	"net/http"
	"testing"
)

func redactHeader(h http.Header) string {
	if h.Get("Authorization") != "" {
		return "Authorization: [redacted]"
	}
	return ""
}

func TestHeaderRedaction(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "Bearer secret")
	s := redactHeader(h)
	if s != "Authorization: [redacted]" {
		t.Fatalf("got %q", s)
	}
}
