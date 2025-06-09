// Copyright 2025 AI Chat
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

package cmd

import (
	"fmt"
	"io"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/spf13/cobra"
)

func newAskCmd(c llm.Client) *cobra.Command {
	var model string
	var temp float64
	var maxTokens int
	cmd := &cobra.Command{
		Use:   "ask <prompt>",
		Args:  cobra.ExactArgs(1),
		Short: "Send a prompt to the LLM",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := llm.Request{
				Model:       model,
				Temperature: temp,
				MaxTokens:   maxTokens,
				Messages:    []llm.Message{{Role: "user", Content: args[0]}},
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
			if _, err := fmt.Fprintln(cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&model, "model", config.GetString("model"), "model name")
	cmd.Flags().Float64Var(&temp, "temperature", 0, "sampling temperature")
	cmd.Flags().IntVar(&maxTokens, "max_tokens", 0, "max tokens")
	return cmd
}
