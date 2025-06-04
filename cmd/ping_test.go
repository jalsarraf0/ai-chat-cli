package cmd

import (
	"bytes"
	"context"
	"testing"
)

type stubClient struct{}

func (stubClient) Ping(ctx context.Context) error              { return nil }
func (stubClient) Version(ctx context.Context) (string, error) { return "", nil }

func TestPingCommand(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	cmd := newPingCmd(stubClient{})
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if got := buf.String(); got != "pong\n" {
		t.Fatalf("expected pong got %q", got)
	}
}

func TestPingCommandVerbose(t *testing.T) {
	t.Setenv("SHELL", "/bin/zsh")
	verbose = true
	defer func() { verbose = false }()
	cmd := newPingCmd(stubClient{})
	out := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SetErr(errBuf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if out.String() != "pong\n" {
		t.Fatalf("expected pong got %q", out.String())
	}
	if errBuf.String() == "" {
		t.Fatalf("expected shell info")
	}
}
