package chat

import "context"

// Client is the interface for interacting with a chat backend.
type Client interface {
	// Prompt sends the given message to the backend and returns the response.
	Prompt(ctx context.Context, message string) (string, error)
}
