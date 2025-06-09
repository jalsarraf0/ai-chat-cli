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

package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAssetsListWriteError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "list"})
	root.SetOut(failWriter{})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestAssetsCatError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "missing.txt"})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestAssetsExportMkdirError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	base := filepath.Join(t.TempDir(), "dir")
	if err := os.MkdirAll(base, 0o755); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(base, "file")
	if err := os.WriteFile(file, []byte{}, 0o600); err != nil {
		t.Fatal(err)
	}
	dest := filepath.Join(file, "out.json")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
