// Copyright 2025 The ai-chat-cli Authors
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
