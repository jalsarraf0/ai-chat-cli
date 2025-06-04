package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompletionCommand(t *testing.T) {
	tests := []struct {
		shell string
		ext   string
	}{
		{"bash", "bash"},
		{"zsh", "zsh"},
		{"fish", "fish"},
		{"powershell", "powershell"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.shell, func(t *testing.T) {
			dir := t.TempDir()
			out := filepath.Join(dir, tt.shell)
			root := NewRootCmd()
			root.SetArgs([]string{"completion", tt.shell, "--out", out})
			if err := root.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if _, err := os.Stat(out); err != nil {
				t.Fatalf("file not created: %v", err)
			}
		})
	}
}
