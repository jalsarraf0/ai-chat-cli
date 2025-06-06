// Copyright 2025 The ai-chat-cli Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cmd provides CLI commands.
package cmd

import (
	"fmt"
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
