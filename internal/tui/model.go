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

// Package tui implements the terminal UI.
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type llmTokenMsg struct {
	Token string
	Done  bool
}

func newSpinner() spinner.Model {
	sp := spinner.New()
	sp.Spinner = spinner.Line
	return sp
}

// Model implements a simple chat interface.
type Model struct {
	history   []string
	viewHist  []string
	cursor    int
	input     textinput.Model
	height    int
	version   string
	light     bool
	styles    Styles
	spinner   spinner.Model
	streaming bool
	streamBuf strings.Builder
	ctrl      Controller
}

// NewModel creates a Model with optional initial rows.
func NewModel(initialRows int) Model {
	ti := textinput.New()
	ti.Prompt = "You: "
	ti.Placeholder = "Ask me anything..."
	ti.Focus()
	ti.CharLimit = 512
	m := Model{input: ti, height: initialRows, spinner: newSpinner()}
	m.styles = LoadStyles("")
	ti.PromptStyle = m.styles.Cursor
	ti.Cursor.Style = m.styles.Cursor
	return m
}

// SetController attaches a Controller for streaming.
func (m *Model) SetController(c Controller) { m.ctrl = c }

// SetVersion sets the application version displayed in the header.
func (m *Model) SetVersion(v string) { m.version = v }

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd { return textinput.Blink }

// UseLightTheme enables the light palette.
func (m *Model) UseLightTheme() {
	m.light = true
	m.styles = LoadStyles("themes/light.json")
	m.input.PromptStyle = m.styles.Cursor
	m.input.Cursor.Style = m.styles.Cursor
}

// UseTheme loads a custom theme by name.
func (m *Model) UseTheme(name string) {
	m.styles = LoadStyles(name)
	m.input.PromptStyle = m.styles.Cursor
	m.input.Cursor.Style = m.styles.Cursor
}

func (m *Model) historyHeight() int {
	if m.height == 0 {
		return 5
	}
	h := m.height - 2 // header + footer
	if h < 5 {
		return 5
	}
	return h
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case llmTokenMsg:
		if msg.Done {
			if m.streamBuf.Len() > 0 {
				m.history = append(m.history, m.streamBuf.String())
			}
			m.viewHist = append([]string{}, m.history...)
			m.streamBuf.Reset()
			m.streaming = false
			return m, nil
		}
		if msg.Token != "" {
			m.streamBuf.WriteString(msg.Token)
			if len(m.viewHist) > len(m.history) {
				m.viewHist[len(m.viewHist)-1] = m.streamBuf.String()
			} else {
				m.viewHist = append(m.viewHist, m.streamBuf.String())
			}
			m.streaming = true
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.input.Width = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyPgUp:
			lines := m.historyHeight()
			if lines > 0 {
				m.cursor += lines
				maxIdx := len(m.viewHist)
				if m.cursor > maxIdx-lines {
					m.cursor = maxIdx - lines
				}
				if m.cursor < 0 {
					m.cursor = 0
				}
			}
		case tea.KeyPgDown:
			lines := m.historyHeight()
			m.cursor -= lines
			if m.cursor < 0 {
				m.cursor = 0
			}
		case tea.KeyEnter:
			val := strings.TrimSpace(m.input.Value())
			if val == ":q" {
				return m, tea.Quit
			}
			if val != "" {
				m.history = append(m.history, val)
				m.viewHist = append(m.viewHist, val)
				m.streamBuf.Reset()
				if m.ctrl.prog != nil {
					m.streaming = true
					m.ctrl.Stream(val)
				}
			}
			m.input.Reset()
			m.cursor = 0
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View renders the UI.
func (m Model) View() string {
	histLines := m.historyHeight()
	start := len(m.viewHist) - histLines - m.cursor
	if start < 0 {
		start = 0
	}
	end := start + histLines
	if end > len(m.viewHist) {
		end = len(m.viewHist)
	}
	lines := m.viewHist[start:end]
	if len(lines) == 0 {
		lines = []string{m.styles.Placeholder.Render("No messages yet – type a prompt and press Enter.")}
	}
	pad := 0
	if len(lines) < histLines {
		pad = (histLines - len(lines)) / 2
	}
	var b strings.Builder
	b.WriteString(strings.Repeat("\n", pad))
	for _, l := range lines {
		b.WriteString(l)
		b.WriteByte('\n')
	}
	historyView := m.styles.History.Height(histLines).Render(strings.TrimRight(b.String(), "\n"))
	inputView := m.styles.Input.Render(m.input.View())
	if m.streaming {
		inputView += " " + m.spinner.View()
	}
	themeName := "dark"
	if m.light {
		themeName = "light"
	}
	header := m.styles.Header.Render(fmt.Sprintf("ai-chat-cli %s [%s]", m.version, themeName))
	footer := m.styles.Footer.Render("Ctrl+C / Esc: quit | PgUp/PgDn: scroll | ↑/↓: history")
	content := lipgloss.JoinVertical(lipgloss.Left, header, historyView, inputView, footer)
	return m.styles.App.Render(content)
}
