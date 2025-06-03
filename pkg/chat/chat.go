package chat

import "context"

// Client abstracts the LLM backend.
type Client interface {
    Ping(ctx context.Context) error
    Version(ctx context.Context) (string, error)
}

type mockClient struct{}

func (mockClient) Ping(ctx context.Context) error { return nil }
func (mockClient) Version(ctx context.Context) (string, error) { return "mock", nil }

// NewMockClient returns a Client for tests.
func NewMockClient() Client { return mockClient{} }
