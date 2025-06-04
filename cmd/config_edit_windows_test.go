//go:build windows

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestConfigEditRunWindows(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	editor := filepath.Join(dir, "ed.bat")
	// simple batch file that exits successfully
	if err := os.WriteFile(editor, []byte("@echo off\r\n"), 0o755); err != nil {
		t.Fatalf("write editor: %v", err)
	}
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	t.Setenv("EDITOR", editor)
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}
