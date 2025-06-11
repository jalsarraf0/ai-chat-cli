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

func TestRootAskError(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	llmClient = errorClient{}
	cmd := newRootCmd()
	cmd.SetArgs([]string{"oops"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
