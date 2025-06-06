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
