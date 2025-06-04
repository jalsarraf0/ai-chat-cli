package tui

import (
	"bytes"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestViewSnapshot(t *testing.T) {
	t.Parallel()
	m := NewModel(10)
	m.height = 10
	m.history = []string{"hello", "world"}
	var tm tea.Model
	tm, _ = m.Update(tea.WindowSizeMsg{Width: 10, Height: 10})
	m = tm.(Model)
	got := m.View()
	want, err := os.ReadFile("testdata/view.golden")
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	if !bytes.Equal([]byte(got), bytes.TrimRight(want, "\n")) {
		t.Fatalf("mismatch\n%s\nvs\n%s", got, want)
	}
}
