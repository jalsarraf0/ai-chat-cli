//go:build !windows

package shell

import "os"

func Detect() Kind {
	return detectFromEnv(os.Getenv("SHELL"), "")
}
