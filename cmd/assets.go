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
	"os"
	"path/filepath"

	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
	"github.com/spf13/cobra"
)

func newAssetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets",
		Short: "Manage embedded assets",
	}
	cmd.AddCommand(newAssetsListCmd())
	cmd.AddCommand(newAssetsCatCmd())
	cmd.AddCommand(newAssetsExportCmd())
	return cmd
}

func newAssetsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List embedded assets",
		RunE: func(cmd *cobra.Command, _ []string) error {
			for _, n := range embedutil.List() {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), n); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func newAssetsCatCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cat <name>",
		Args:  cobra.ExactArgs(1),
		Short: "Print asset content",
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := embedutil.Read(args[0])
			if err != nil {
				return err
			}
			_, err = cmd.OutOrStdout().Write(b)
			return err
		},
	}
}

func newAssetsExportCmd() *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "export <name> <file>",
		Args:  cobra.ExactArgs(2),
		Short: "Export asset to file",
		RunE: func(_ *cobra.Command, args []string) error {
			data, err := embedutil.Read(args[0])
			if err != nil {
				return err
			}
			dest := args[1]
			if !force {
				if _, err := os.Stat(dest); err == nil {
					return fmt.Errorf("%s exists", dest)
				} else if !os.IsNotExist(err) {
					return err
				}
			}
			if err := os.MkdirAll(filepath.Dir(dest), 0o750); err != nil {
				return err
			}
			return os.WriteFile(dest, data, 0o600)
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite destination")
	return cmd
}
