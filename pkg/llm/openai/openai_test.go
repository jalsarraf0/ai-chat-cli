//go:build integration

package openai

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func TestLiveCompletion(t *testing.T) {
	if os.Getenv("AI_CHAT_API_KEY") == "" {
		t.Skip("AI_CHAT_API_KEY not set")
	}
	c := New()
	s, err := c.Completion(context.Background(), llm.Request{
		Model:    "gpt-3.5-turbo",
		Messages: []llm.Message{{Role: "user", Content: "ping"}},
	})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	_, err = s.Recv()
	if err != nil && err != io.EOF {
		t.Fatalf("recv: %v", err)
	}
}
