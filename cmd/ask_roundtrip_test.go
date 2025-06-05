//go:build unit

package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCommand(t *testing.T) {
	out := new(bytes.Buffer)
	root := newRootCmd()
	root.SetOut(out)
	root.SetArgs([]string{"ask", "hi"})
	root.SilenceErrors = true
	root.SilenceUsage = true
	root.PersistentPreRun = func(*cobra.Command, []string) {}
	if err := newAskCmd(mock.New("ok")).RunE(root, []string{"hi"}); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("got %q", out.String())
	}
}
