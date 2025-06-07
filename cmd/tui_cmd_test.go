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
	root.SetArgs([]string{"tui", "--theme", "light", "--height", "5"})
	root.SetIn(bytes.NewBufferString(":q\n"))
	root.SetOut(io.Discard)
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !ran || height != 5 {
		t.Fatalf("flags not set or program not run")
	}
}
