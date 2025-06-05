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
