package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model implements a simple chat interface.
type Model struct {
	history []string // chat log
	cursor  int      // scroll offset
	input   textinput.Model
	height  int
	light   bool
	styles  Styles
}

// NewModel creates a Model with optional initial rows.
func NewModel(initialRows int) Model {
	ti := textinput.New()
	ti.Prompt = "› "
	m := Model{input: ti, height: initialRows}
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
	case tea.WindowSizeMsg:
		if m.height == 0 && msg.Height > 0 {
			m.height = msg.Height
		} else {
			m.height = msg.Height
		}
		m.input.Width = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyPgUp:
			lines := m.historyHeight()
			if lines > 0 {
				m.cursor += lines
				max := len(m.history)
				if m.cursor > max-lines {
					m.cursor = max - lines
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
	return lipgloss.JoinVertical(lipgloss.Left, historyView, inputView)
}
