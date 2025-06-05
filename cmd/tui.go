package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jalsarraf0/ai-chat-cli/internal/tui"
	"github.com/spf13/cobra"
)

var teaRun = func(p *tea.Program) (tea.Model, error) { return p.Run() }

var light bool
var height int

func newTuiCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tui", Short: "Interactive terminal UI"}
	cmd.Flags().BoolVar(&light, "light", false, "use light theme")
	cmd.Flags().IntVar(&height, "height", 0, "override initial rows (0 = auto)")
	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		m := tui.NewModel(height)
		if light {
			m.UseLightTheme()
		}
		p := tea.NewProgram(m)
		_, err := teaRun(p)
		return err
	}
	return cmd
}
