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
