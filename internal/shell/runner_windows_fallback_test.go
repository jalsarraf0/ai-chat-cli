//go:build windows

package shell

import (
	"context"
	"os/exec"
	"testing"
)

// TestRunWindowsFallback verifies cmd is used when PowerShell is absent.
func TestRunWindowsFallback(t *testing.T) {
	t.Parallel()
	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(name string) (string, error) {
		if name == "cmd" {
			return exec.LookPath("cmd")
		}
		return "", exec.ErrNotFound
	}
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()
	if _, _, err := Run(context.Background(), "echo", "hi"); err != nil {
		t.Fatalf("run: %v", err)
	}
}
