//go:build unit

package cmd

import (
	"bytes"
	"io"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// fakeProgram is a stub implementing Run.
type fakeProgram struct{}

func (fakeProgram) Run() (tea.Model, error) { return nil, nil }

func TestTuiCmd(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	ran := false
	teaRun = func(p *tea.Program) (tea.Model, error) { ran = true; return nil, nil }
	defer func() { teaRun = func(p *tea.Program) (tea.Model, error) { return p.Run() } }()

	root := newRootCmd()
	root.SetArgs([]string{"tui", "--light", "--height", "5"})
	root.SetIn(bytes.NewBufferString(":q\n"))
	root.SetOut(io.Discard)
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !ran || !light || height != 5 {
		t.Fatalf("flags not set or program not run")
	}
}
