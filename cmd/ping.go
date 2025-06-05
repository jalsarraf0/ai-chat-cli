// Package cmd provides CLI commands.
package cmd

import (
	"context"
	"fmt"

	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
func newPingCmd(c chat.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check connectivity",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := c.Ping(context.Background()); err != nil {
				return err
			}
			if verbose {
				if _, err := fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s\n", detectedShell); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), "pong"); err != nil {
				return err
			}
			return nil
		},
	}
}
