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

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "config", Short: "Manage configuration"}
	cmd.AddCommand(newConfigShowCmd())
	cmd.AddCommand(newConfigSetCmd())
	cmd.AddCommand(newConfigGetCmd())
	cmd.AddCommand(newConfigListCmd())
	cmd.AddCommand(newConfigEditCmd())
	return cmd
}

func newConfigShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show configuration",
		RunE: func(cmd *cobra.Command, _ []string) error {
			p := configPath()
			b, err := os.ReadFile(p) // #nosec G304 -- path from user input
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(cmd.OutOrStdout(), "%s\n%s", p, string(b))
			return err
		},
	}
}

func newConfigSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration key",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			return config.Set(args[0], args[1])
		},
	}
}

func newConfigGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !config.IsSet(args[0]) {
				return fmt.Errorf("key not found")
			}
			_, err := fmt.Fprintln(cmd.OutOrStdout(), config.Get(args[0]))
			return err
		},
	}
}

func newConfigListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List configuration",
		RunE: func(cmd *cobra.Command, _ []string) error {
			for k, v := range config.All() {
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s: %v\n", k, v); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func newConfigEditCmd() *cobra.Command {
	var dryRun bool
	c := &cobra.Command{
		Use:   "edit",
		Short: "Edit configuration file",
		RunE: func(_ *cobra.Command, _ []string) error {
			if dryRun {
				return nil
			}
			editor := os.Getenv("EDITOR")
			if editor == "" {
				if runtime.GOOS == "windows" {
					editor = "notepad"
				} else {
					editor = "vi"
				}
			}
			e := exec.Command(editor, configPath()) // #nosec G204 -- user-controlled editor
			e.Stdin = os.Stdin
			e.Stdout = os.Stdout
			e.Stderr = os.Stderr
			return e.Run()
		},
	}
	c.Flags().BoolVar(&dryRun, "dry-run", false, "skip launching editor")
	return c
}

func configPath() string {
	return config.Path()
}
