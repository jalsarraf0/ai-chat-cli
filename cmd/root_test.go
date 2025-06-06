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

package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestRootExecute(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		want   string
		stderr bool
	}{
		{"ping", []string{"ping"}, "pong\n", false},
		{"ping-verbose", []string{"--verbose", "ping"}, "pong\n", true},
		{"version", []string{"version", "--short"}, "1.2.3\n", false},
	}
	Version = "1.2.3"
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", "/bin/bash")
			t.Setenv("AICHAT_OPENAI_API_KEY", "key")
			config.Reset()
			cfg := filepath.Join(t.TempDir(), "c.yaml")
			cmd := newRootCmd()
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			cmd.SetOut(outBuf)
			cmd.SetErr(errBuf)
			cmd.SetArgs(append([]string{"--config", cfg}, tt.args...))
			if err := cmd.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if outBuf.String() != tt.want {
				t.Fatalf("want %q got %q", tt.want, outBuf.String())
			}
			if tt.stderr && errBuf.String() == "" {
				t.Fatalf("expected stderr output")
			}
		})
	}
}

func TestExecuteFailure(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		os.Args = []string{"ai-chat", "ping"}
		Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExecuteFailure")
	env := []string{"GO_WANT_HELPER_PROCESS=1"}
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, "AICHAT_OPENAI_API_KEY=") {
			env = append(env, e)
		}
	}
	cmd.Env = env
	// Ensure config.Load fails by not setting AICHAT_OPENAI_API_KEY
	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected exit")
	}
	if e, ok := err.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Fatalf("unexpected error: %v", err)
	}
}
