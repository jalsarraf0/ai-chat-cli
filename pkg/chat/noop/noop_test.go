package noop

import (
	"context"
	"testing"
)

func TestClientPrompt(t *testing.T) {
	t.Parallel()
	var c Client
	resp, err := c.Prompt(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Prompt error: %v", err)
	}
	if resp != "stub" {
		t.Fatalf("expected stub, got %q", resp)
	}
}
