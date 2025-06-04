package mock

import (
	"context"
	"io"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func TestStream(t *testing.T) {
	t.Parallel()
	c := New("hi", "there")
	s, err := c.Completion(context.Background(), llm.Request{})
	if err != nil {
		t.Fatalf("completion: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "hi" {
		t.Fatalf("first token: %v %v", r, err)
	}
	r, err = s.Recv()
	if err != nil || r.Content != "there" {
		t.Fatalf("second token: %v %v", r, err)
	}
	_, err = s.Recv()
	if err != io.EOF {
		t.Fatalf("want EOF")
	}
}
