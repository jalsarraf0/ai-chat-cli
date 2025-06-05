//go:build !windows

package shell

import "os"

// Detect returns the current shell on Unix.
func Detect() Kind {
	return detectFromEnv(os.Getenv("SHELL"), "")
}
