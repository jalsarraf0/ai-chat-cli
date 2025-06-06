//go:build unit

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWatchCmd(t *testing.T) {
	cfgDir := t.TempDir()
	cfgFile := filepath.Join(cfgDir, "watch.yaml")
	if err := os.WriteFile(cfgFile, []byte("patterns:\n  - foo"), 0o600); err != nil {
		t.Fatalf("write cfg: %v", err)
	}
	cmd := newWatchCmd()
	cmd.SetArgs([]string{"--config", cfgFile})
	cmd.SetIn(strings.NewReader("foo\nbar\nfoo bar\n"))
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("want 2 alerts got %d", len(lines))
	}
}
