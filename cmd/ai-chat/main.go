package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	// Version is injected by build flags.
	Version string
	// Commit is injected by build flags.
	Commit string
	// BuiltAt is injected by build flags.
	BuiltAt string
)

func run() error {
	return newRootCmd().Execute()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var cfgFile string
	var debug bool

	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "ai-chat is a CLI interface to chat systems",
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug output")

	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newPingCmd())
	cmd.AddCommand(newCompletionCmd(cmd))
	cmd.AddCommand(newManCmd(cmd))

	return cmd
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print build information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Version: %s\nCommit: %s\nBuiltAt: %s\n", Version, Commit, BuiltAt)
		},
	}
}

func newPingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check connectivity",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "pong")
		},
	}
}

func newCompletionCmd(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:       "completion [bash|zsh|fish|ps1]",
		Short:     "Generate completion script",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"bash", "zsh", "fish", "ps1"},
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return root.GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				return root.GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				return root.GenFishCompletion(cmd.OutOrStdout(), true)
			case "ps1":
				return root.GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
			default:
				return fmt.Errorf("unknown shell %s", args[0])
			}
		},
	}
}

func newManCmd(root *cobra.Command) *cobra.Command {
	var dir string
	cmd := &cobra.Command{
		Use:   "man",
		Short: "Generate man pages",
		RunE: func(cmd *cobra.Command, args []string) error {
			header := &doc.GenManHeader{
				Title:   "AI-CHAT",
				Section: "1",
			}
			return doc.GenManTree(root, header, dir)
		},
	}
	cmd.Flags().StringVar(&dir, "dir", "dist/man", "directory to write man pages")
	return cmd
}
