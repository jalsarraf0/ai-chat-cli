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
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
        Version = "v1.0.6"
	Commit  = "none"
	Date    = "unknown"
)

func newVersionCmd(version, commit, date string) *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, _ []string) {
			if short {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), version); err != nil {
					cmd.Println("Error:", err)
				}
				return
			}
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s %s %s\n", version, commit, date); err != nil {
				cmd.Println("Error:", err)
			}
		},
	}
	cmd.Flags().BoolVarP(&short, "short", "s", false, "short output")
	return cmd
}
