package cmd

import (
	"bytes"
	"testing"
)

func TestModelsCmd(t *testing.T) {
	cmd := newModelsCmd(stubLLM{models: []string{"b", "a"}})
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if out.String() != "a\nb\n" {
		t.Fatalf("out=%q", out.String())
	}
}

func TestModelsCmdError(t *testing.T) {
	cmd := newModelsCmd(errLLM{})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
