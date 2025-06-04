package shell

import (
	"context"
	"errors"
	"testing"
)

func TestRunCmd_LookPathStub(t *testing.T) {

	orig := lookPath
	defer func() { lookPath = orig }()

	// Stub lookPath to return a fake binary.
	lookPath = func(file string) (string, error) {
		if file == "echo" {
			return "/bin/echo", nil
		}
		return "", errors.New("unexpected binary")
	}

	stdout, stderr, err := runCmd(context.Background(), "echo", []string{"hi"})
	if err != nil {
		t.Fatalf("runCmd failed: %v", err)
	}
	if stdout != "hi\n" || stderr != "" {
		t.Fatalf("unexpected output: %q / %q", stdout, stderr)
	}
}
