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

package tui

import (
	"os"

	"github.com/charmbracelet/lipgloss"

	"github.com/jalsarraf0/ai-chat-cli/pkg/theme"
)

// Styles defines UI colours.
type Styles struct {
	App     lipgloss.Style
	History lipgloss.Style
	Input   lipgloss.Style
	Cursor  lipgloss.Style
}

// LoadStyles builds styles from the embedded theme.
func LoadStyles(name string) Styles {
	p := theme.Load(name)
	app := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(p.Background)).Padding(0, 1)
	hist := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(p.Background)).Height(0)
	inp := lipgloss.NewStyle().BorderTop(true).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color(p.Background))
	cursor := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffff"))
	if os.Getenv("NO_COLOR") != "" {
		app = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(0, 1)
		hist = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Height(0)
		inp = lipgloss.NewStyle().BorderTop(true).BorderStyle(lipgloss.NormalBorder())
		cursor = lipgloss.NewStyle()
	}
	return Styles{App: app, History: hist, Input: inp, Cursor: cursor}
}
