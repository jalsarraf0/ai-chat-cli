package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "AI Chat CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "ai-chat version %s (%s)\n", version, commit)
			return err
		},
	}
	cmd.SetVersionTemplate("ai-chat version {{.Version}} (" + commit + ")\n")
	cmd.Version = version
	return cmd
}

var run = func() error { return newRootCmd().Execute() }

func realMain() int {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

func main() { os.Exit(realMain()) }
