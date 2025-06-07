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
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
)

func TestAssetsCommands(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	out := new(bytes.Buffer)
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "list"})
	if err := root.Execute(); err != nil {
		t.Fatalf("list: %v", err)
	}
	lines := bytes.Split(bytes.TrimSpace(out.Bytes()), []byte("\n"))
	if len(lines) != 4 {
		t.Fatalf("want 4 lines got %d", len(lines))
	}

	// cat
	root = newRootCmd()
	out.Reset()
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "templates/system.tmpl"})
	if err := root.Execute(); err != nil {
		t.Fatalf("cat: %v", err)
	}
	want, _ := embedutil.Read("templates/system.tmpl")
	if !bytes.Equal(out.Bytes(), want) {
		t.Fatalf("cat output mismatch")
	}

	// export
	dest := filepath.Join(t.TempDir(), "x", "y.json")
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err != nil {
		t.Fatalf("export: %v", err)
	}
	b1, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("read dest: %v", err)
	}
	b2, _ := embedutil.Read("themes/light.json")
	if !bytes.Equal(b1, b2) {
		t.Fatalf("export mismatch")
	}
	// overwrite with --force
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest, "--force"})
	if err := root.Execute(); err != nil {
		t.Fatalf("export force: %v", err)
	}
}
