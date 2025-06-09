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

package main_test

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
	if out, err := exec.Command("go", "build", "-o", exe, "-ldflags", ldflags, "./cmd/ai-chat").CombinedOutput(); err != nil {
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
