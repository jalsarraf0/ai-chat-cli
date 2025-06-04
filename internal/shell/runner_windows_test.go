//go:build windows

package shell

import (
	"context"
	"testing"
)

func TestRunWindows(t *testing.T) {
	orig := lookPath
	lookPath = func(string) (string, error) { return "pwsh", nil }
	defer func() { lookPath = orig }()
	out, errOut, err := Run(context.Background(), "Write-Output", "ok")
	if err != nil {
		t.Fatalf("run err %v", err)
	}
	if out == "" || errOut != "" {
		t.Fatalf("unexpected out %q err %q", out, errOut)
	}
}
