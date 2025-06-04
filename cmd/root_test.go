package cmd

import (
	"bytes"
	"testing"
)

func TestRootExecutePing(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	cmd := newRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"ping"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got := buf.String(); got != "pong\n" {
		t.Fatalf("got %q", got)
	}
}

func TestRootExecutePingVerbose(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	cmd := newRootCmd()
	cmd.SetArgs([]string{"--verbose", "ping"})
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.SetOut(outBuf)
	cmd.SetErr(errBuf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	t.Logf("stdout=%q stderr=%q", outBuf.String(), errBuf.String())
	if outBuf.String() != "pong\n" || errBuf.String() == "" {
		t.Fatalf("unexpected output")
	}
}

func TestRootExecuteVersion(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	cmd := newRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"version", "--short"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if buf.String() == "" {
		t.Fatalf("expected version")
	}
}
