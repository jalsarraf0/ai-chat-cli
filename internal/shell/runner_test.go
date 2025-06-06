// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
