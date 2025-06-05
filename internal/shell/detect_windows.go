//go:build windows

package shell

import "os"

// Detect returns the current shell on Windows.
func Detect() Kind {
	if k := detectFromEnv(os.Getenv("SHELL"), os.Getenv("ComSpec")); k != Unknown {
		return k
	}
	return Cmd
}
