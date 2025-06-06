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

package cmd

import (
	"bytes"
	"testing"

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
