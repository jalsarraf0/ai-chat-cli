
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
	"context"
	"fmt"
	"sort"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/spf13/cobra"
)

func newModelsCmd(c llm.Client) *cobra.Command {
	var list bool
	cmd := &cobra.Command{
		Use:   "models",
		Short: "List available models",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !list {
				list = true
			}
			if list {
				models, err := c.ListModels(context.Background())
				if err != nil {
					return err
				}
				sort.Strings(models)
				for _, m := range models {
					if _, err := fmt.Fprintln(cmd.OutOrStdout(), m); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&list, "list", false, "list models")
	return cmd
}
