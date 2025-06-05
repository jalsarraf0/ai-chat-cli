// Package chat defines the LLM chat interface.
package chat

import "context"

// Client abstracts the LLM backend.
type Client interface {
	Ping(ctx context.Context) error
	Version(ctx context.Context) (string, error)
}

type mockClient struct{}

func (mockClient) Ping(_ context.Context) error              { return nil }
func (mockClient) Version(_ context.Context) (string, error) { return "mock", nil }

// NewMockClient returns a Client for tests.
func NewMockClient() Client { return mockClient{} }
