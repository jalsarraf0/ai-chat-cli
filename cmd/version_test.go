package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"short", []string{"--short"}, "1.2.3\n"},
		{"full", nil, "1.2.3 abc now\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", "/bin/bash")
			cmd := newVersionCmd("1.2.3", "abc", "now")
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetArgs(tt.args)
			if err := cmd.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if buf.String() != tt.want {
				t.Fatalf("want %q got %q", tt.want, buf.String())
			}
		})
	}
}

type versionErrWriter struct{}

func (versionErrWriter) Write(_ []byte) (int, error) { return 0, fmt.Errorf("werr") }

func TestVersionWriteError(t *testing.T) {
	cmd := newVersionCmd("1.2.3", "abc", "now")
	cmd.SetOut(versionErrWriter{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
