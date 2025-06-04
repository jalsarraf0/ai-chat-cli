package cmd

import (
	"bytes"
	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"os"
	"path/filepath"
	"runtime"
	"testing"
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
	if out.Len() == 0 {
		t.Fatalf("expected output")
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
