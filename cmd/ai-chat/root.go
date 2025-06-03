package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// newRootCmd creates the root command for ai-chat.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "Interact with AI chat models",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), "ai-chat CLI bootstrap"); err != nil {
				return fmt.Errorf("print greeting: %w", err)
			}
			return nil
		},
	}
	return cmd
}

// Execute runs the ai-chat CLI.
func Execute(ctx context.Context) error {
	return newRootCmd().ExecuteContext(ctx)
}
