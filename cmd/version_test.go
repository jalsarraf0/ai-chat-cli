package cmd

import (
	"bytes"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	t.Parallel()
	cmd := newVersionCmd("1.0.0", "abc", "now")
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	want := "1.0.0 abc now\n"
	if out.String() != want {
		t.Fatalf("want %q got %q", want, out.String())
	}
	out.Reset()
	cmd.SetArgs([]string{"--short"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("short: %v", err)
	}
	if out.String() != "1.0.0\n" {
		t.Fatalf("short wrong")
	}
}
