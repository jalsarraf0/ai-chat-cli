package cmd

import (
	"fmt"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/spf13/cobra"
)

// Version is set via -ldflags.
var Version string

// NewRootCmd returns the root command for the CLI.
func NewRootCmd() *cobra.Command {
	var debug bool
	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "AI Chat CLI",
	}

	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug output")

	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), Version)
			return nil
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "ping",
		Short: "Test shell",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			out, errOut, err := shell.Run(ctx, "echo pong")
			if debug {
				name, path, derr := shell.Detect()
				if derr == nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s path=%s\n", name, path)
				}
			}
			if err != nil {
				cmd.ErrOrStderr().Write(errOut)
				return err
			}
			cmd.OutOrStdout().Write(out)
			return nil
		},
	})

	return cmd
}
