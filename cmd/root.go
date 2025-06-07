// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package cmd provides CLI commands.
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	chatClient    chat.Client = chat.NewMockClient()
	llmClient     llm.Client  = mock.New("hello")
	verbose       bool
	detectedShell shell.Kind
	cfgFile       string
)

func newRootCmd() *cobra.Command {
	detectedShell = shell.Detect()
	cmd := &cobra.Command{
		Use:   "ai-chat",
		Short: "Interact with AI chat services",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := config.Load(cfgFile); err != nil {
				return err
			}
			log.Printf("INFO: config %s", config.Path())
			if verbose {
				if _, err := fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s\n", detectedShell); err != nil {
					cmd.Println("Error:", err)
				}
			}
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default auto)")
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	cmd.AddCommand(newPingCmd(chatClient))
	cmd.AddCommand(newVersionCmd(Version, Commit, Date))
	cmd.AddCommand(newAssetsCmd())
	cmd.AddCommand(newConfigCmd())
	cmd.AddCommand(newTuiCmd())
	cmd.AddCommand(newAskCmd(llmClient))
	cmd.AddCommand(newAIOpsCmd())
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
