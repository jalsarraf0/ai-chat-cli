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

// Package tui implements the terminal UI.
package tui

import (
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
	history   []string // chat log
	cursor    int      // scroll offset
	input     textinput.Model
	height    int
	light     bool
	styles    Styles
	spinner   spinner.Model
	streaming bool
}

// NewModel creates a Model with optional initial rows.
func NewModel(initialRows int) Model {
	ti := textinput.New()
	ti.Prompt = "› "
	m := Model{input: ti, height: initialRows, spinner: newSpinner()}
	m.styles = LoadStyles(false)
	ti.PromptStyle = m.styles.Cursor
	ti.Cursor.Style = m.styles.Cursor
	return m
}

// Init satisfies tea.Model.
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// UseLightTheme enables the light palette.
func (m *Model) UseLightTheme() {
	m.light = true
	m.styles = LoadStyles(true)
	m.input.PromptStyle = m.styles.Cursor
	m.input.Cursor.Style = m.styles.Cursor
}

func (m *Model) historyHeight() int {
	if m.height == 0 {
		return 0
	}
	return m.height * 80 / 100
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case llmTokenMsg:
		if msg.Token != "" {
			m.history = append(m.history, msg.Token)
		}
		if msg.Done {
			m.streaming = false
		} else {
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
				maxIdx := len(m.history)
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
	start := len(m.history) - histLines - m.cursor
	if start < 0 {
		start = 0
	}
	end := start + histLines
	if end > len(m.history) {
		end = len(m.history)
	}
	var b strings.Builder
	for i := start; i < end; i++ {
		b.WriteString(m.history[i])
		b.WriteByte('\n')
	}
	historyView := m.styles.History.Render(b.String())
	inputView := m.styles.Input.Render(m.input.View())
	if m.streaming {
		inputView += " " + m.spinner.View()
	}
	return lipgloss.JoinVertical(lipgloss.Left, historyView, inputView)
}
