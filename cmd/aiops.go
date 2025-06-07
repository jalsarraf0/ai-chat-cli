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
	"bufio"
	"fmt"
	"os"

	"github.com/jalsarraf0/ai-chat-cli/internal/aiops"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type watchConfig struct {
	Patterns []string `yaml:"patterns"`
}

func newAIOpsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "aiops", Short: "AI-Ops helper"}
	cmd.AddCommand(newWatchCmd())
	return cmd
}

func newWatchCmd() *cobra.Command {
	var cfgPath string
	cmd := &cobra.Command{
		Use:   "watch --config <file>",
		Short: "Stream logs from stdin and detect anomalies",
		RunE: func(cmd *cobra.Command, _ []string) error {
			data, err := os.ReadFile(cfgPath) // #nosec G304 -- file from user input
			if err != nil {
				return err
			}
			var cfg watchConfig
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return err
			}
			det, err := aiops.NewRegexDetector(cfg.Patterns)
			if err != nil {
				return err
			}
			scanner := bufio.NewScanner(cmd.InOrStdin())
			for scanner.Scan() {
				alert, ok := det.Detect(cmd.Context(), scanner.Text())
				if ok {
					if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", alert.Pattern, alert.Line); err != nil {
						return err
					}
				}
			}
			return scanner.Err()
		},
	}
	cmd.Flags().StringVar(&cfgPath, "config", "", "config file")
	if err := cmd.MarkFlagRequired("config"); err != nil {
		cmd.PrintErrln(err)
	}
	return cmd
}
