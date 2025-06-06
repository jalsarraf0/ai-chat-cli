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
