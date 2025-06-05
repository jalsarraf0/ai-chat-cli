//go:build windows

package shell

import (
	"context"
)

// Run executes a command via the user's shell on Windows.
func Run(ctx context.Context, cmd string, args ...string) (string, string, error) {
	lookPathMu.RLock()
	lp := lookPath
	lookPathMu.RUnlock()
	pwsh, _ := lp("pwsh")
	if pwsh == "" {
		pwsh, _ = lp("powershell")
	}
	if pwsh != "" {
		return runCmd(ctx, pwsh, append([]string{"-NoProfile", "-Command", cmd}, args...))
	}
	return runCmd(ctx, "cmd", append([]string{"/C", cmd}, args...))
}
