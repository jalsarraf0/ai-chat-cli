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
