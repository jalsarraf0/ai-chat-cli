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
			if tt.name == "ping" || tt.name == "ping-verbose" || tt.name == "version" {
				t.Setenv("OPENAI_API_KEY", "")
			} else {
				t.Setenv("OPENAI_API_KEY", "key")
			}
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
		os.Args = []string{"ai-chat", "hello"}
		Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExecuteFailure")
	env := []string{"GO_WANT_HELPER_PROCESS=1"}
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, "OPENAI_API_KEY=") {
			env = append(env, e)
		}
	}
	cmd.Env = env
	// Ensure config.Load fails by not setting OPENAI_API_KEY
	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected exit")
	}
	if e, ok := err.(*exec.ExitError); !ok || e.ExitCode() == 0 {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRootModels(t *testing.T) {
	llmClient = stubLLM{models: []string{"m1"}}
	t.Setenv("OPENAI_API_KEY", "")
	config.Reset()
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	cmd := newRootCmd()
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SetArgs([]string{"--config", cfg, "models"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("models: %v", err)
	}
	if out.String() != "m1\n" {
		t.Fatalf("out=%q", out.String())
	}
}
