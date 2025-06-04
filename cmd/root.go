/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	chatClient    chat.Client = chat.NewMockClient()
	verbose       bool
	detectedShell shell.Kind
)

func NewRootCmd() *cobra.Command {
	var cfgFile string
	detectedShell = shell.Detect()
	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "Interact with AI chat services",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				if _, err := fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s\n", detectedShell); err != nil {
					cmd.Println("Error:", err)
				}
			}
		},
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cmd.AddCommand(newPingCmd(chatClient))
	cmd.AddCommand(newVersionCmd(Version, Commit, Date))
	cmd.AddCommand(newCompletionCmd(cmd))
	cmd.AddCommand(newInitCmd(cmd))
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
