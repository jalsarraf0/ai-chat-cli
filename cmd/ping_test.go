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
