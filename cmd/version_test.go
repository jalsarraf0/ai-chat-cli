package cmd

import (
    "bytes"
    "testing"
)

func TestVersionShort(t *testing.T) {
    t.Parallel()
    cmd := newVersionCmd("1.2.3", "abc", "now")
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)
    cmd.SetArgs([]string{"--short"})
    if err := cmd.Execute(); err != nil {
        t.Fatalf("execute: %v", err)
    }
    if got := buf.String(); got != "1.2.3\n" {
        t.Fatalf("expected version got %q", got)
    }
}

func TestVersionFull(t *testing.T) {
    t.Parallel()
    cmd := newVersionCmd("1.2.3", "abc", "now")
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)
    cmd.SetArgs(nil)
    if err := cmd.Execute(); err != nil {
        t.Fatalf("execute: %v", err)
    }
    if got := buf.String(); got != "1.2.3 abc now\n" {
        t.Fatalf("expected full version got %q", got)
    }
}
