package cmd

import (
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestExecute(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExecute", "--")
	env := append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	if d := os.Getenv("GOCOVERDIR"); d != "" {
		env = append(env, "GOCOVERDIR="+d)
	}
	cmd.Env = env
	if err := cmd.Run(); err != nil {
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			return
		}
		t.Fatalf("run: %v", err)
	}
}

func TestExecuteInProc(t *testing.T) {
	oldArgs := os.Args
	os.Args = []string{"ai-chat", "version", "--short"}
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	Version = "test"
	Execute()
	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}
	os.Stdout = old
	out, _ := io.ReadAll(r)
	os.Args = oldArgs
	if string(out) != "test\n" {
		t.Fatalf("want test got %q", string(out))
	}
}

func TestExecuteError(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_ERROR_PROCESS") == "1" {
		chatClient = errClient{}
		os.Args = []string{"ai-chat", "ping"}
		Execute()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExecuteError", "--")
	env := append(os.Environ(), "GO_WANT_HELPER_ERROR_PROCESS=1")
	if d := os.Getenv("GOCOVERDIR"); d != "" {
		env = append(env, "GOCOVERDIR="+d)
	}
	cmd.Env = env
	if err := cmd.Run(); err == nil {
		t.Fatalf("expected exit error")
	}
}
