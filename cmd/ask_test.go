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

func TestAskCmd(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	c := mock.New("h", "i")
	cmd := newAskCmd(c)
	cmd.SetArgs([]string{"hello"})
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if buf.String() != "hi\n" {
		t.Fatalf("output %q", buf.String())
	}
}

func TestAskCmdFlags(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	c := mock.New("z")
	cmd := newAskCmd(c)
	cmd.SetArgs([]string{"--model", "gpt-4", "--temperature", "0.5", "--max_tokens", "1", "ping"})
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

func TestAskCmdError(t *testing.T) {
	t.Parallel()
	cmd := newAskCmd(errorClient{})
	cmd.SetArgs([]string{"oops"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
