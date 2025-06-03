/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
        "os"

        "github.com/spf13/cobra"
        "github.com/jalsarraf0/ai-chat-cli/pkg/chat"
)



// rootCmd represents the base command when called without any subcommands
var chatClient chat.Client = chat.NewMockClient()

func newRootCmd() *cobra.Command {
        var (
                cfgFile string
                verbose bool
        )
        cmd := &cobra.Command{
                Use:   "ai-chat",
                Short: "Interact with AI chat services",
        }
        cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
        cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
        cmd.AddCommand(newPingCmd(chatClient))
        cmd.AddCommand(newVersionCmd(Version, Commit, Date))
        return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := newRootCmd().Execute(); err != nil {
                os.Exit(1)
        }
}


