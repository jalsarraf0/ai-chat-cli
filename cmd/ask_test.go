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
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestRootAsk(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	buf := new(bytes.Buffer)
	llmClient = mock.New("h", "i")
	cmd := newRootCmd()
	cmd.SetArgs([]string{"hello"})
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if buf.String() != "hi\n" {
		t.Fatalf("output %q", buf.String())
	}
}

func TestRootAskFlags(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	buf := new(bytes.Buffer)
	llmClient = mock.New("z")
	cmd := newRootCmd()
	cmd.SetArgs([]string{"--model", "gpt-4", "--temperature", "0.5", "--max-tokens", "1", "hello"})
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if buf.String() != "z\n" {
		t.Fatalf("output %q", buf.String())
	}
}

type errorClient struct{}

func (errorClient) Completion(_ context.Context, _ llm.Request) (llm.Stream, error) {
	return nil, io.EOF
}
func (errorClient) ListModels(context.Context) ([]string, error) { return nil, nil }

func TestRootAskError(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	llmClient = errorClient{}
	cmd := newRootCmd()
	cmd.SetArgs([]string{"oops"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestRootAskStdin(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	buf := new(bytes.Buffer)
	llmClient = mock.New("ok")
	cmd := newRootCmd()
	cmd.SetIn(bytes.NewBufferString("hello"))
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if buf.String() != "ok\n" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestRootAskArgAndStdin(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	llmClient = mock.New("done")
	buf := new(bytes.Buffer)
	cmd := newRootCmd()
	cmd.SetArgs([]string{"question"})
	cmd.SetIn(bytes.NewBufferString("context"))
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if buf.String() != "done\n" {
		t.Fatalf("got %q", buf.String())
	}
}

func TestRootAskFileStdin(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	dir := t.TempDir()
	filePath := filepath.Join(dir, "in.txt")
	if err := os.WriteFile(filePath, []byte("file"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	f, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	t.Cleanup(func() { _ = f.Close() })
	llmClient = mock.New("f")
	out := new(bytes.Buffer)
	cmd := newRootCmd()
	cmd.SetIn(f)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if out.String() != "f\n" {
		t.Fatalf("got %q", out.String())
	}
}

func TestRootAskHelp(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	llmClient = mock.New()
	out := new(bytes.Buffer)
	cmd := newRootCmd()
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if !strings.Contains(out.String(), "Interact with AI chat services") {
		t.Fatalf("unexpected %q", out.String())
	}
}
