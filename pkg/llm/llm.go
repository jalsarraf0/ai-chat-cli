package llm

import "context"

// Message represents a chat message.
type Message struct {
	Role    string
	Content string
}

// Request describes a completion query.
type Request struct {
	Model       string
	Temperature float64
	MaxTokens   int
	Messages    []Message
}

// Response is a single chunk of completion text.
type Response struct {
	Content string
}

// Stream yields completion tokens.
type Stream interface {
	Recv() (Response, error)
}

// Client sends completion requests to a language model.
type Client interface {
	Completion(ctx context.Context, req Request) (Stream, error)
}
