package cmd

import (
    "bytes"
    "testing"
)

func TestRootExecutePing(t *testing.T) {
    t.Parallel()
    cmd := newRootCmd()
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)
    cmd.SetArgs([]string{"ping"})
    if err := cmd.Execute(); err != nil {
        t.Fatalf("execute: %v", err)
    }
    if got := buf.String(); got != "pong\n" {
        t.Fatalf("got %q", got)
    }
}

func TestRootExecuteVersion(t *testing.T) {
    t.Parallel()
    cmd := newRootCmd()
    buf := new(bytes.Buffer)
    cmd.SetOut(buf)
    cmd.SetArgs([]string{"version", "--short"})
    if err := cmd.Execute(); err != nil {
        t.Fatalf("execute: %v", err)
    }
    if buf.String() == "" {
        t.Fatalf("expected version")
    }
}
