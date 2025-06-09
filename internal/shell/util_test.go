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

package shell

import "testing"

func TestClassify(t *testing.T) {
	cases := map[string]Kind{
		"bash":     Bash,
		"zsh":      Zsh,
		"fish":     Fish,
		"pwsh.exe": PowerShell,
		"cmd":      Cmd,
		"unknown":  Unknown,
	}
	for in, want := range cases {
		if got := classify(in); got != want {
			t.Fatalf("%s => %v want %v", in, got, want)
		}
	}
}

func TestKindString(t *testing.T) {
	if Bash.String() != "bash" {
		t.Fatalf("String() mismatch")
	}
}

func TestDetectFromEnv(t *testing.T) {
	got := detectFromEnv("/bin/zsh", "")
	if got != Zsh {
		t.Fatalf("want zsh got %v", got)
	}
	if detectFromEnv("", `C:/Windows/System32/cmd.exe`) != Cmd {
		t.Fatalf("cmd detection failed")
	}
}
