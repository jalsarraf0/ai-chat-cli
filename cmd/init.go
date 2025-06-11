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
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var setupScript = "./setup.sh"

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Run interactive setup wizard",
		RunE: func(cmd *cobra.Command, args []string) error {
			script := setupScript
			if !filepath.IsAbs(script) {
				exe, err := os.Executable()
				if err == nil {
					dir := filepath.Dir(exe)
					if _, err := os.Stat(filepath.Join(dir, script)); err == nil {
						script = filepath.Join(dir, script)
					}
				}
			}
			w := exec.Command(script, args...)
			w.Stdin = os.Stdin
			w.Stdout = cmd.OutOrStdout()
			w.Stderr = cmd.ErrOrStderr()
			return w.Run()
		},
	}
}
