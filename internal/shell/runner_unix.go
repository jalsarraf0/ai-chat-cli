//go:build !windows

package shell

import (
	"context"
	"os"
)

// Run executes a command via the user's shell on Unix.
func Run(ctx context.Context, cmd string, args ...string) (string, string, error) {
	sh := os.Getenv("SHELL")
	if sh == "" {
		sh = "/bin/sh"
	}
	all := append([]string{"-c", cmd}, args...)
	return runCmd(ctx, sh, all)
}
