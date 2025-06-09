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
	"bufio"
	"fmt"
	"os"

	"github.com/jalsarraf0/ai-chat-cli/internal/aiops"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
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
