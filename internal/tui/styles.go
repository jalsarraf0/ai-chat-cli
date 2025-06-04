package tui

import (
	"encoding/json"

	"github.com/charmbracelet/lipgloss"
	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
)

type palette struct {
	Background string `json:"background"`
}

// Styles defines layout styling.
type Styles struct {
	History lipgloss.Style
	Input   lipgloss.Style
	Cursor  lipgloss.Style
}

// LoadStyles builds styles from the embedded theme.
func LoadStyles(light bool) Styles {
	name := "themes/dark.json"
	if light {
		name = "themes/light.json"
	}
	data, err := embedutil.Read(name)
	if err != nil {
		data = []byte(`{"background":""}`)
	}
	var p palette
	_ = json.Unmarshal(data, &p)
	hist := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(p.Background)).Height(0)
	inp := lipgloss.NewStyle().BorderTop(true).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(p.Background))
	cursor := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffff"))
	return Styles{History: hist, Input: inp, Cursor: cursor}
}
