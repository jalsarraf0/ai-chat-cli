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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWatchCmd(t *testing.T) {
	cfgDir := t.TempDir()
	cfgFile := filepath.Join(cfgDir, "watch.yaml")
	if err := os.WriteFile(cfgFile, []byte("patterns:\n  - foo"), 0o600); err != nil {
		t.Fatalf("write cfg: %v", err)
	}
	cmd := newWatchCmd()
	cmd.SetArgs([]string{"--config", cfgFile})
	cmd.SetIn(strings.NewReader("foo\nbar\nfoo bar\n"))
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("want 2 alerts got %d", len(lines))
	}
}
