// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/openai"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	chatClient    chat.Client = chat.NewMockClient()
	defaultLLM    llm.Client  = mock.New("hello")
	llmClient     llm.Client  = defaultLLM
	verbose       bool
	detectedShell shell.Kind
	cfgFile       string
	model         string
	temp          float64
	maxTokens     int
	stream        bool
	profile       string
)

func skipCfgValidation(cmd *cobra.Command) bool {
	root := cmd.Root().Name()
	prefixes := []string{
		root + " config",
		root + " version",
		root + " ping",
		root + " assets",
		root + " completion",
		root + " healthcheck",
		root + " login",
	}
	for _, p := range prefixes {
		if strings.HasPrefix(cmd.CommandPath(), p) {
			return true
		}
	}
	return false
}

func newRootCmd() *cobra.Command {
	detectedShell = shell.Detect()
	cmd := &cobra.Command{
		Use:   "ai-chat [prompt]",
		Short: "Interact with AI chat services",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return askRunE(llmClient)(cmd, args)
		},
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if skipCfgValidation(cmd) {
				config.SkipValidation(true)
				defer config.SkipValidation(false)
			}
			if err := config.Load(cfgFile); err != nil {
				if errors.Is(err, config.ErrAPIKeyMissing) {
					return fmt.Errorf("%w\nset with 'ai-chat login <key>' or env OPENAI_API_KEY", err)
				}
				return err
			}
			if reflect.DeepEqual(llmClient, defaultLLM) {
				llmClient = openai.New()
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
	cmd.PersistentFlags().StringVarP(&model, "model", "m", config.GetString("model"), "model name")
	cmd.PersistentFlags().Float64VarP(&temp, "temperature", "t", 0, "sampling temperature")
	cmd.PersistentFlags().IntVar(&maxTokens, "max-tokens", 0, "max tokens")
	cmd.PersistentFlags().BoolVar(&stream, "stream", false, "stream partial deltas")
	cmd.PersistentFlags().StringVar(&profile, "profile", "", "named config profile")
	cmd.AddCommand(newPingCmd(chatClient))
	cmd.AddCommand(newVersionCmd(Version, Commit, Date))
	cmd.AddCommand(newAssetsCmd())
	cmd.AddCommand(newConfigCmd())
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newModelsCmd(llmClient))
	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newWizardCmd())
	cmd.AddCommand(newTuiCmd())
	cmd.AddCommand(newAskCmd())
	cmd.AddCommand(newHealthcheckCmd())
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

func askRunE(c llm.Client) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var prompt string
		if len(args) > 0 {
			prompt = args[0]
		}
		in := cmd.InOrStdin()
		if f, ok := in.(*os.File); ok {
			if st, err := f.Stat(); err == nil && (st.Mode()&os.ModeCharDevice) == 0 {
				data, err := io.ReadAll(in)
				if err != nil {
					return err
				}
				if prompt == "" {
					prompt = string(data)
				} else {
					prompt += "\n" + string(data)
				}
			}
		} else {
			if data, err := io.ReadAll(in); err == nil && len(data) > 0 {
				if prompt == "" {
					prompt = string(data)
				} else {
					prompt += "\n" + string(data)
				}
			}
		}
		if prompt == "" {
			return cmd.Help()
		}
		req := llm.Request{
			Model:       model,
			Temperature: temp,
			MaxTokens:   maxTokens,
			Messages:    []llm.Message{{Role: "user", Content: prompt}},
		}
		stream, err := c.Completion(cmd.Context(), req)
		if err != nil {
			return err
		}
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			if _, err := fmt.Fprint(cmd.OutOrStdout(), resp.Content); err != nil {
				return err
			}
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout())
		return err
	}
}
