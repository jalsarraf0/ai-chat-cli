package cmd

import (
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestLoginCmd(t *testing.T) {
	dir := t.TempDir()
	cfg := filepath.Join(dir, "c.yaml")
	config.Reset()
	c := newRootCmd()
	c.SetArgs([]string{"--config", cfg, "login", "abc123"})
	if err := c.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if config.GetString("openai_api_key") != "abc123" {
		t.Fatalf("key not set")
	}
}
