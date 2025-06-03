package shell

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrUnsupported is returned when the current shell cannot be determined.
var ErrUnsupported = errors.New("unsupported shell")

// Detect returns the shell name and executable path based on environment variables.
// It checks PSExecutionPolicy (PowerShell), ComSpec (cmd.exe), and SHELL (POSIX).
func Detect() (string, string, error) {
	if os.Getenv("PSExecutionPolicy") != "" {
		path := os.Getenv("SHELL")
		if path == "" {
			// Fallback to generic powershell name; may rely on PATH.
			path = "powershell"
		}
		return "powershell", path, nil
	}
	if cs := os.Getenv("ComSpec"); cs != "" {
		return "cmd", cs, nil
	}
	if sh := os.Getenv("SHELL"); sh != "" {
		return filepath.Base(sh), sh, nil
	}
	return "", "", ErrUnsupported
}
