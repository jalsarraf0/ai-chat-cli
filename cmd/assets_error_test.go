//go:build unit

package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAssetsExportErrors(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	dest := filepath.Join(t.TempDir(), "out.json")

	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "bogus", dest})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}

	// existing file without --force
	if err := os.WriteFile(dest, []byte("x"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	err := root.Execute()
	if err == nil || !strings.Contains(err.Error(), "exists") {
		t.Fatalf("want exists error, got %v", err)
	}
}
