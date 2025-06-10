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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jalsarraf0/ai-chat-cli/internal/tui"
	"github.com/spf13/cobra"
)

var teaRun = func(p *tea.Program) (tea.Model, error) { return p.Run() }

var (
	themeFlag string
	height    int
)

func newTuiCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tui", Short: "Interactive terminal UI"}
	cmd.Flags().StringVar(&themeFlag, "theme", "", "theme name (light/dark)")
	cmd.Flags().IntVar(&height, "height", 0, "override initial rows (0 = auto)")
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		m := tui.NewModel(height)
		if themeFlag == "light" {
			m.UseLightTheme()
		} else if themeFlag != "" {
			m.UseTheme(themeFlag)
		}
		p := tea.NewProgram(m)
		_, err := teaRun(p)
		return err
	}
	return cmd
}
