package cmd

import (
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
