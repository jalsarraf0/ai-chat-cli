package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
)

type stubClient struct{}

type errClient struct{}

func (stubClient) Ping(ctx context.Context) error              { return nil }
func (stubClient) Version(ctx context.Context) (string, error) { return "", nil }
func (errClient) Ping(ctx context.Context) error               { return context.DeadlineExceeded }
func (errClient) Version(ctx context.Context) (string, error)  { return "", context.DeadlineExceeded }

func TestPingCommand(t *testing.T) {
	tests := []struct {
		name    string
		shell   string
		verbose bool
		client  chat.Client
		wantErr bool
	}{
		{"normal", "/bin/bash", false, stubClient{}, false},
		{"verbose", "/bin/zsh", true, stubClient{}, false},
		{"error", "/bin/bash", false, errClient{}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", tt.shell)
			verbose = tt.verbose
			defer func() { verbose = false }()
			cmd := newPingCmd(tt.client)
			out := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(errBuf)
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if !tt.wantErr && out.String() != "pong\n" {
				t.Fatalf("out=%q", out.String())
			}
			if tt.verbose && errBuf.String() == "" {
				t.Fatalf("expected shell info")
			}
		})
	}
}
