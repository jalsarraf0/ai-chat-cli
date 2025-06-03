package main

import (
	"bytes"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	t.Parallel()
	Version = "v0.0.0-test"
	cmd := newVersionCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	out := buf.String()
	if !bytes.Contains(buf.Bytes(), []byte(Version)) {
		t.Fatalf("expected version in output: %s", out)
	}
}

func TestPingCmd(t *testing.T) {
	t.Parallel()
	cmd := newPingCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got := buf.String(); got != "pong\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestRootCommands(t *testing.T) {
	t.Parallel()
	cmd := newRootCmd()
	cmd.SetArgs([]string{"ping"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("root ping: %v", err)
	}

	cmd = newRootCmd()
	cmd.SetArgs([]string{"completion", "bash"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("root completion: %v", err)
	}

	shells := []string{"zsh", "fish", "ps1"}
	for _, sh := range shells {
		cmd = newRootCmd()
		cmd.SetArgs([]string{"completion", sh})
		if err := cmd.Execute(); err != nil {
			t.Fatalf("root completion %s: %v", sh, err)
		}
	}

	dir := t.TempDir()
	cmd = newRootCmd()
	cmd.SetArgs([]string{"man", "--dir", dir})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("root man: %v", err)
	}

	// Ensure run executes without error for a simple command.
	if err := run(); err != nil {
		t.Fatalf("run error: %v", err)
	}
}
