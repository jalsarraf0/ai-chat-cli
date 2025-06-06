//go:build unit

package cmd

import (
	"bytes"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCmdFlags(t *testing.T) {
	out := new(bytes.Buffer)
	cmd := newAskCmd(mock.New("ok"))
	cmd.SetOut(out)
	cmd.SetArgs([]string{"--model", "gpt-4", "--temperature", "0.5", "--max_tokens", "5", "hi"})
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("unexpected %q", out.String())
	}
}
