package cmd

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"ping"})
	if err := root.ExecuteContext(context.Background()); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if buf.String() != "pong\n" {
		t.Errorf("got %q", buf.String())
	}
}

func TestPingDebug(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	root := NewRootCmd()
	out := &bytes.Buffer{}
	errb := &bytes.Buffer{}
	root.SetOut(out)
	root.SetErr(errb)
	root.SetArgs([]string{"--debug", "ping"})
	if err := root.ExecuteContext(context.Background()); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if out.String() != "pong\n" {
		t.Errorf("out %q", out.String())
	}
	if !bytes.Contains(errb.Bytes(), []byte("shell=")) {
		t.Errorf("err %q", errb.String())
	}
}

func TestVersion(t *testing.T) {
	root := NewRootCmd()
	Version = "v1.2.3"
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"version"})
	if err := root.ExecuteContext(context.Background()); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if strings.TrimSpace(buf.String()) != Version {
		t.Errorf("version %q", buf.String())
	}
}

func TestPingError(t *testing.T) {
	// unsupported shell triggers an error path
	t.Setenv("SHELL", "")
	root := NewRootCmd()
	root.SetArgs([]string{"ping"})
	if err := root.ExecuteContext(context.Background()); err == nil {
		t.Fatalf("expected error")
	}
}
