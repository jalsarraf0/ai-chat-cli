package noop

import (
	"context"

	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
)

// Client implements chat.Client returning a stub response.
type Client struct{}

// Prompt returns a static stub response.
func (Client) Prompt(ctx context.Context, message string) (string, error) {
	return "stub", nil
}

// Ensure Client implements chat.Client.
var _ chat.Client = Client{}
