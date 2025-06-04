//go:build windows

package shell

import "testing"

func TestDetectWindows(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("ComSpec", `C:\\Windows\\System32\\cmd.exe`)
	if got := Detect(); got != Cmd {
		t.Fatalf("want cmd got %v", got)
	}
	t.Setenv("SHELL", `C:\\Program Files\\PowerShell\\pwsh.exe`)
	if got := Detect(); got != PowerShell {
		t.Fatalf("want pwsh got %v", got)
	}
}
