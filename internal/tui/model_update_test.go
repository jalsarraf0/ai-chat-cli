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
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestModelUpdate(t *testing.T) {
	t.Parallel()
	m := NewModel(10)
	m.input.SetValue("msg")
	var cmd tea.Cmd
	var tm tea.Model
	tm, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = tm.(Model)
	if cmd != nil {
		t.Fatalf("unexpected cmd %v", cmd)
	}
	if len(m.history) != 1 {
		t.Fatalf("expected 1 line")
	}
	m.height = 10
	m.history = append(m.history, make([]string, 20)...)
	tm, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgUp})
	m = tm.(Model)
	if m.cursor == 0 {
		t.Fatalf("pgup failed")
	}
	old := m.cursor
	tm, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
	m = tm.(Model)
	if m.cursor >= old {
		t.Fatalf("pgdn failed")
	}
	_, quit := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if quit == nil {
		t.Fatalf("expect quit")
	}
}
func TestQuitSequence(t *testing.T) {
	t.Parallel()
	m := NewModel(5)
	m.input.SetValue(":q")
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatalf("expected quit")
	}
}

func TestWindowResize(t *testing.T) {
	t.Parallel()
	m := NewModel(0)
	tm, _ := m.Update(tea.WindowSizeMsg{Width: 40, Height: 20})
	m = tm.(Model)
	if m.historyHeight() != 16 {
		t.Fatalf("height calc")
	}
}
func TestUseLightTheme(t *testing.T) {
	t.Parallel()
	m := NewModel(5)
	m.UseLightTheme()
	if !m.light {
		t.Fatalf("flag not set")
	}
}
func TestInitCmd(t *testing.T) {
	t.Parallel()
	m := NewModel(0)
	if cmd := m.Init(); cmd == nil {
		t.Fatal("no command returned")
	}
}

func TestLoadStylesDiff(t *testing.T) {
	t.Parallel()
	_ = LoadStyles(false)
	_ = LoadStyles(true)
}
func TestHistoryHeightZero(t *testing.T) {
	t.Parallel()
	m := NewModel(0)
	if m.historyHeight() != 0 {
		t.Fatalf("want 0")
	}
}

func TestQuitEsc(t *testing.T) {
	t.Parallel()
	m := NewModel(5)
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if cmd == nil {
		t.Fatalf("expect quit")
	}
}
func TestNoOpMessage(t *testing.T) {
	t.Parallel()
	m := NewModel(5)
	_, _ = m.Update(struct{}{})
}
func TestScrollBounds(t *testing.T) {
	t.Parallel()
	m := NewModel(10)
	m.height = 10
	for i := 0; i < 5; i++ {
		m.history = append(m.history, "x")
	}
	m.cursor = 100
	tm, _ := m.Update(tea.KeyMsg{Type: tea.KeyPgUp})
	m = tm.(Model)
	if m.cursor < 0 {
		t.Fatalf("invalid cursor")
	}
	tm, _ = m.Update(tea.KeyMsg{Type: tea.KeyPgDown})
	m = tm.(Model)
	if m.cursor != 0 {
		t.Fatalf("scroll reset")
	}
}
