//go:build windows

package shell

import (
	"context"
	"testing"
)

func TestRunWindows(t *testing.T) {
	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(string) (string, error) { return "pwsh", nil }
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()
	out, errOut, err := Run(context.Background(), "Write-Output", "ok")
	if err != nil {
		t.Fatalf("run err %v", err)
	}
	if out == "" || errOut != "" {
		t.Fatalf("unexpected out %q err %q", out, errOut)
	}
}
