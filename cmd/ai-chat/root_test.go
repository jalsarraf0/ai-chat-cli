package main

import (
	"bytes"
	"context"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Parallel()

	buf := new(bytes.Buffer)
	cmd := newRootCmd()
	cmd.SetOut(buf)
	if err := cmd.ExecuteContext(context.Background()); err != nil {
		t.Fatalf("execute: %v", err)
	}
	got := buf.String()
	want := "ai-chat CLI bootstrap\n"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestExecuteFunction(t *testing.T) {
	t.Parallel()
	if err := Execute(context.Background()); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}
}
