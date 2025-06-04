package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestRootExecute(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		want   string
		stderr bool
	}{
		{"ping", []string{"ping"}, "pong\n", false},
		{"ping-verbose", []string{"--verbose", "ping"}, "pong\n", true},
		{"version", []string{"version", "--short"}, "1.2.3\n", false},
	}
	Version = "1.2.3"
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", "/bin/bash")
			t.Setenv("AICHAT_OPENAI_API_KEY", "key")
			config.Reset()
			cfg := filepath.Join(t.TempDir(), "c.yaml")
			cmd := newRootCmd()
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			cmd.SetOut(outBuf)
			cmd.SetErr(errBuf)
			cmd.SetArgs(append([]string{"--config", cfg}, tt.args...))
			if err := cmd.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if outBuf.String() != tt.want {
				t.Fatalf("want %q got %q", tt.want, outBuf.String())
			}
			if tt.stderr && errBuf.String() == "" {
				t.Fatalf("expected stderr output")
			}
		})
	}
}
