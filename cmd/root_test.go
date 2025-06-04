package cmd

import (
	"bytes"
	"testing"
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
			cmd := NewRootCmd()
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			cmd.SetOut(outBuf)
			cmd.SetErr(errBuf)
			cmd.SetArgs(tt.args)
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
