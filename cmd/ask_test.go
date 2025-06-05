package cmd

import (
	"bytes"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm/mock"
)

func TestAskCmd(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	c := mock.New("h", "i")
	cmd := newAskCmd(c)
	cmd.SetArgs([]string{"hello"})
	cmd.SetOut(buf)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if buf.String() != "hi\n" {
		t.Fatalf("output %q", buf.String())
	}
}
