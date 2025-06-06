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

package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestStampedVersion(t *testing.T) {
	t.Parallel()
	tmp := t.TempDir()
	exe := filepath.Join(tmp, "ai-chat")
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	ldflags := "-X github.com/jalsarraf0/ai-chat-cli/cmd.Version=1.0.0 -X github.com/jalsarraf0/ai-chat-cli/cmd.Commit=abc123 -X github.com/jalsarraf0/ai-chat-cli/cmd.Date=now"
	if out, err := exec.Command("go", "build", "-o", exe, "-ldflags", ldflags, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	cmd := exec.Command(exe, "version")
	cmd.Env = append(cmd.Env, "AICHAT_OPENAI_API_KEY=dummy")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if !strings.Contains(string(out), "1.0.0 abc123 now") {
		t.Fatalf("unexpected output: %s", out)
	}
}
