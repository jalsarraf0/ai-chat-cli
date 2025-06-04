package cmd

import (
	"bytes"
	"io"
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
			b, err := os.ReadFile(out)
			if err != nil {
				t.Fatalf("read: %v", err)
			}
			if !bytes.HasPrefix(b, []byte("#")) {
				t.Fatalf("invalid completion")
			}
		})
	}
	t.Run("readonly", func(t *testing.T) {
		file := filepath.Join(t.TempDir(), "f")
		if err := os.WriteFile(file, []byte{}, 0644); err != nil {
			t.Fatalf("write: %v", err)
		}
		out := filepath.Join(file, "x")
		root := NewRootCmd()
		root.SetArgs([]string{"completion", "bash", "--out", out})
		if err := root.Execute(); err == nil {
			t.Fatalf("expected error")
		}
	})
	t.Run("invalid", func(t *testing.T) {
		root := NewRootCmd()
		root.SetArgs([]string{"completion", "csh", "--out", filepath.Join(t.TempDir(), "csh")})
		if err := root.Execute(); err == nil {
			t.Fatalf("expected error")
		}
	})
}

func TestGenerateCompletionInvalid(t *testing.T) {
	if err := generateCompletion(NewRootCmd(), "csh", io.Discard); err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateCompletionAll(t *testing.T) {
	shells := []string{"bash", "zsh", "fish", "powershell"}
	for _, sh := range shells {
		if err := generateCompletion(NewRootCmd(), sh, io.Discard); err != nil {
			t.Fatalf("%s: %v", sh, err)
		}
	}
}

func TestCompletionDefaultOut(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	defer func() {
		if err := os.Chdir(old); err != nil {
			t.Fatalf("chdir back: %v", err)
		}
	}()
	root := NewRootCmd()
	root.SetArgs([]string{"completion", "bash"})
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if _, err := os.Stat(filepath.Join("dist", "completion", "bash")); err != nil {
		t.Fatalf("default out: %v", err)
	}
}
