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
