package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
)

func TestInitDryRun(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	root.SetArgs([]string{"init", "--dry-run"})
	out := new(bytes.Buffer)
	root.SetOut(out)
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if !bytes.Contains(out.Bytes(), []byte("✅")) {
		t.Fatalf("missing success message")
	}
	if _, err := os.Stat(filepath.Join(home, ".ai-chat", "completion.bash")); err == nil {
		t.Fatalf("completion file created in dry-run")
	}
}

func TestInitNonDryRun(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	root.SetArgs([]string{"init", "--no-prompt"})
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if _, err := os.Stat(filepath.Join(home, ".ai-chat", "completion.bash")); err != nil {
		t.Fatalf("completion not created: %v", err)
	}
}

func TestInitHelpers(t *testing.T) {
	if rcPath("/home", shell.Bash) != "/home/.bashrc" {
		t.Fatalf("rcPath bash")
	}
	rcPath("/home", shell.Zsh)
	rcPath("/home", shell.Fish)
	rcPath("/home", shell.PowerShell)
	if sourceLine(shell.PowerShell, "f") != ". \"f\"\n" {
		t.Fatalf("sourceLine ps1")
	}
	if promptFor(shell.Fish) == "" || promptFor(shell.Bash) == "" || promptFor(shell.PowerShell) == "" {
		t.Fatalf("promptFor")
	}
	f, err := openAppend(filepath.Join(t.TempDir(), "x"), 0o644)
	if err != nil {
		t.Fatalf("openAppend: %v", err)
	}
	f.Close()
}
