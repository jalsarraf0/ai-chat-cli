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

func TestInitWithPrompt(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	root.SetArgs([]string{"init"})
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(home, ".bashrc"))
	if err != nil {
		t.Fatalf("read rc: %v", err)
	}
	if !bytes.Contains(data, []byte("(ai)›")) {
		t.Fatalf("prompt not added")
	}
}

func TestInitNoPrompt(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	root.SetArgs([]string{"init", "--no-prompt"})
	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(home, ".bashrc"))
	if err != nil {
		t.Fatalf("read rc: %v", err)
	}
	if !bytes.Contains(data, []byte("source")) {
		t.Fatalf("rc missing source")
	}
}

func TestInitHelpers(t *testing.T) {
	t.Parallel()
	if rcPath("/home", shell.Bash) != "/home/.bashrc" {
		t.Fatalf("rcPath bash")
	}
	if rcPath("/home", shell.Zsh) != "/home/.zshrc" {
		t.Fatalf("rcPath zsh")
	}
	if rcPath("/home", shell.Fish) != "/home/.config/fish/config.fish" {
		t.Fatalf("rcPath fish")
	}
	if rcPath("/home", shell.PowerShell) != "/home/Documents/WindowsPowerShell/profile.ps1" {
		t.Fatalf("rcPath ps")
	}
	if rcPath("/home", shell.Kind("x")) != "" {
		t.Fatalf("rcPath unknown")
	}
	if sourceLine(shell.PowerShell, "f") != ". \"f\"\n" {
		t.Fatalf("sourceLine ps1")
	}
	if promptFor(shell.Fish) == "" || promptFor(shell.Bash) == "" || promptFor(shell.PowerShell) == "" || promptFor(shell.Zsh) == "" || promptFor(shell.Kind("x")) != "" {
		t.Fatalf("promptFor")
	}
	f, err := openAppend(filepath.Join(t.TempDir(), "x"), 0o644)
	if err != nil {
		t.Fatalf("openAppend: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	if _, err := openAppend(filepath.Join("/no", "way", "file"), 0o644); err == nil {
		t.Fatalf("expected open error")
	}
}

func TestInitUnsupportedShell(t *testing.T) {
	t.Setenv("SHELL", "/bin/unknown")
	root := NewRootCmd()
	root.SetArgs([]string{"init"})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
