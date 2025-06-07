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

//go:build unit

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAssetsRoundTrip(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	dest := filepath.Join(t.TempDir(), "out.json")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err != nil {
		t.Fatalf("export: %v", err)
	}
	root = newRootCmd()
	out := new(bytes.Buffer)
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "themes/light.json"})
	if err := root.Execute(); err != nil {
		t.Fatalf("cat: %v", err)
	}
	b1, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !bytes.Equal(b1, out.Bytes()) {
		t.Fatalf("round trip mismatch")
	}
}
