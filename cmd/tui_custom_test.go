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

package cmd

import (
	"bytes"
	"io"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTuiCmdCustomTheme(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	ran := false
	teaRun = func(_ *tea.Program) (tea.Model, error) { ran = true; return nil, nil }
	defer func() { teaRun = func(p *tea.Program) (tea.Model, error) { return p.Run() } }()

	root := newRootCmd()
	root.SetArgs([]string{"tui", "--theme", "themes/dark.json", "--height", "1"})
	root.SetIn(bytes.NewBufferString(":q\n"))
	root.SetOut(io.Discard)
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !ran {
		t.Fatalf("program not run")
	}
}
