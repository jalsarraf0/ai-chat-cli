package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestMainVersion(t *testing.T) {
	t.Parallel()
	cmd := newRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"--version"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute failed: %v", err)
	}
	if buf.Len() == 0 {
		t.Fatalf("no output")
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	if err := run(); err != nil {
		t.Fatalf("run failed: %v", err)
	}
}

func TestRealMain(t *testing.T) {
	orig := run
	if code := realMain(); code != 0 {
		t.Fatalf("unexpected exit code %d", code)
	}
	run = func() error { return errors.New("fail") }
	if code := realMain(); code == 0 {
		t.Fatalf("expected failure code")
	}
	run = orig
}
