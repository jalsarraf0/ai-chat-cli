package cmd

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestConfigCommands(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
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
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "show"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestConfigSetInvalid(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "set", "model", "bad"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
