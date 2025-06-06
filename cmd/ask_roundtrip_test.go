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
	"github.com/spf13/cobra"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCommand(t *testing.T) {
	out := new(bytes.Buffer)
	root := newRootCmd()
	root.SetOut(out)
	root.SetArgs([]string{"ask", "hi"})
	root.SilenceErrors = true
	root.SilenceUsage = true
	root.PersistentPreRun = func(*cobra.Command, []string) {}
	if err := newAskCmd(mock.New("ok")).RunE(root, []string{"hi"}); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("got %q", out.String())
	}
}
