//go:build unit

package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWatchCmd_BadConfig(t *testing.T) {
	dir := t.TempDir()
	cfg := filepath.Join(dir, "bad.yaml")
	if err := os.WriteFile(cfg, []byte(":\n"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	cmd := newWatchCmd()
	cmd.SetArgs([]string{"--config", cfg})
	cmd.SetIn(strings.NewReader(""))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
