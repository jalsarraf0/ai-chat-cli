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
	m.SetVersion("v1.0.2")
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
