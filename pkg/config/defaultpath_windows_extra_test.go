//go:build windows

package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDefaultPathTemp ensures fallback to temp when APPDATA and HOME are empty.
func TestDefaultPathTemp(t *testing.T) {
	t.Parallel()
	Reset()
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("APPDATA", "")
	t.Setenv("HOME", "")
	p := defaultPath()
	want := filepath.Join(os.TempDir(), "ai-chat", "config.yaml")
	if p != want {
		t.Fatalf("want %s got %s", want, p)
	}
}
