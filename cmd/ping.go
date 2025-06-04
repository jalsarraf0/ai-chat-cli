/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Ping(context.Background()); err != nil {
				return err
			}
			if verbose {
				fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s\n", detectedShell)
			}
			fmt.Fprintln(cmd.OutOrStdout(), "pong")
			return nil
		},
	}
}
