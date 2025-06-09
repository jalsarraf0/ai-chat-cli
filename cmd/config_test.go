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
	"runtime"
	"strings"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestConfigCommands(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	config.Reset()
	cfg := newRootCmd()
	cfg.SetArgs([]string{"--config", dir + "/c.yaml", "config", "set", "openai_api_key", "abc"})
	if err := cfg.Execute(); err != nil {
		t.Fatalf("set: %v", err)
	}
	show := newRootCmd()
	out := new(bytes.Buffer)
	show.SetOut(out)
	show.SetArgs([]string{"--config", dir + "/c.yaml", "config", "show"})
	if err := show.Execute(); err != nil {
		t.Fatalf("show: %v", err)
	}
	outStr := out.String()
	want := dir + string(os.PathSeparator) + "c.yaml"
	want = filepath.ToSlash(want)
	if !strings.Contains(filepath.ToSlash(outStr), want) {
		t.Fatalf("missing path in output: %s", outStr)
	}
	edit := newRootCmd()
	edit.SetArgs([]string{"--config", dir + "/c.yaml", "config", "edit", "--dry-run"})
	if err := edit.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}

func TestConfigShowMissing(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "show"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestConfigSetInvalid(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "set", "model", "bad"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestConfigEditRun(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
	}
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	editor := filepath.Join(dir, "ed.sh")
	if err := os.WriteFile(editor, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("write editor: %v", err)
	}
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	t.Setenv("EDITOR", editor)
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}

func TestConfigEditDefaultEditor(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
	}
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	vi := filepath.Join(dir, "vi")
	if err := os.WriteFile(vi, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("write vi: %v", err)
	}
	t.Setenv("PATH", dir+":"+os.Getenv("PATH"))
	t.Setenv("EDITOR", "")
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}
