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
