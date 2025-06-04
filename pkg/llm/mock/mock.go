package mock

import (
	"context"
	"io"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

// Client returns predetermined tokens.
type Client struct {
	tokens []string
}

// New creates a mock client that streams the given tokens.
func New(tokens ...string) Client {
	return Client{tokens: tokens}
}

func (c Client) Completion(ctx context.Context, req llm.Request) (llm.Stream, error) {
	return &stream{tokens: c.tokens}, nil
}

type stream struct {
	tokens []string
	i      int
}

func (s *stream) Recv() (llm.Response, error) {
	if s.i >= len(s.tokens) {
		return llm.Response{}, io.EOF
	}
	t := s.tokens[s.i]
	s.i++
	return llm.Response{Content: t}, nil
}
