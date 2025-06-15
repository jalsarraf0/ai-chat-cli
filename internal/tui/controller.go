package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jalsarraf0/ai-chat-cli/internal/openai"
)

// streamCmd streams OpenAI responses as llmTokenMsg.
func streamCmd(prompt string) tea.Cmd {
	return func() tea.Msg {
		ch := make(chan tea.Msg)
		go func() {
			for token := range openai.Stream(prompt) {
				ch <- llmTokenMsg{Token: token, Done: false}
			}
			ch <- llmTokenMsg{Done: true}
			close(ch)
		}()
		return <-ch
	}
}
