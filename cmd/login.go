package cmd

import (
	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/spf13/cobra"
)

func newLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login <openai_api_key>",
		Short: "Set OpenAI API key",
		Args:  cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return config.Set("openai_api_key", args[0])
		},
	}
}
