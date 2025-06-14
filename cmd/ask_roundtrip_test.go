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
	"testing"

	"github.com/spf13/cobra"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCommand(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	out := new(bytes.Buffer)
	llmClient = mock.New("ok")
	root := newRootCmd()
	root.SetOut(out)
	root.SetArgs([]string{"ask", "hi"})
	root.SilenceErrors = true
	root.SilenceUsage = true
	root.PersistentPreRun = func(*cobra.Command, []string) {}
	if err := root.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("got %q", out.String())
	}
}
