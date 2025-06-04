package shell

import (
	"context"
	"errors"
	"os/exec"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	t.Parallel()
	out, errOut, err := Run(context.Background(), "echo", "ok")
	if err != nil {
		t.Fatalf("run err %v", err)
	}
	if out == "" || errOut != "" {
		t.Fatalf("unexpected out=%q err=%q", out, errOut)
	}
}

func TestRunError(t *testing.T) {
	t.Parallel()
	_, _, err := Run(context.Background(), "false")
	if err == nil {
		t.Fatalf("expected error")
	}
	var ee *exec.ExitError
	if !errors.As(err, &ee) {
		t.Fatalf("want ExitError got %T", err)
	}
}

func TestRunDefaultShell(t *testing.T) {
	t.Setenv("SHELL", "")
	out, _, err := Run(context.Background(), "echo", "hi")
	if err != nil || out == "" {
		t.Fatalf("unexpected: %v %q", err, out)
	}
}

func TestRunCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := Run(ctx, "echo", "hi"); err == nil {
		t.Fatalf("expected error")
	}
}

func TestRunCmd(t *testing.T) {
	out, errOut, err := runCmd(context.Background(), "sh", []string{"-c", "echo hi"})
	if err != nil || out != "hi\n" || errOut != "" {
		t.Fatalf("unexpected %v %q %q", err, out, errOut)
	}
}
