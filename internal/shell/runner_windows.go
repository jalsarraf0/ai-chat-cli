//go:build windows

package shell

import "context"

func Run(ctx context.Context, cmd string, args ...string) (string, string, error) {
	pwsh, _ := lookPath("pwsh")
	if pwsh == "" {
		pwsh, _ = lookPath("powershell")
	}
	if pwsh != "" {
		return runCmd(ctx, pwsh, append([]string{"-NoProfile", "-Command", cmd}, args...))
	}
	return runCmd(ctx, "cmd", append([]string{"/C", cmd}, args...))
}
