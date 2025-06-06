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
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCmdFlags(t *testing.T) {
	out := new(bytes.Buffer)
	cmd := newAskCmd(mock.New("ok"))
	cmd.SetOut(out)
	cmd.SetArgs([]string{"--model", "gpt-4", "--temperature", "0.5", "--max_tokens", "5", "hi"})
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("unexpected %q", out.String())
	}
}
