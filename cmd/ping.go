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
	"context"
	"fmt"

	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
func newPingCmd(c chat.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check connectivity",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := c.Ping(context.Background()); err != nil {
				return err
			}
			if verbose {
				if _, err := fmt.Fprintf(cmd.ErrOrStderr(), "shell=%s\n", detectedShell); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), "pong"); err != nil {
				return err
			}
			return nil
		},
	}
}
