package cmd

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
)

type errClient struct{}

func (errClient) Ping(context.Context) error              { return errors.New("fail") }
func (errClient) Version(context.Context) (string, error) { return "", nil }

func TestPingCommand(t *testing.T) {
	t.Parallel()
	// success
	cmd := newPingCmd(chat.NewMockClient())
	out := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.SetOut(out)
	cmd.SetErr(errBuf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("ping success: %v", err)
	}
	if out.String() != "pong\n" {
		t.Fatalf("want pong got %q", out.String())
	}
	if errBuf.Len() != 0 {
		t.Fatalf("unexpected stderr output")
	}
	// failure
	cmd = newPingCmd(errClient{})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}

	verbose = true
	detectedShell = shell.Bash
	cmd = newPingCmd(chat.NewMockClient())
	out.Reset()
	errBuf.Reset()
	cmd.SetOut(out)
	cmd.SetErr(errBuf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("verbose: %v", err)
	}
	verbose = false
	if errBuf.String() == "" {
		t.Fatalf("expected stderr in verbose mode")
	}
}
