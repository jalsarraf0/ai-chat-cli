//go:build windows

package shell

import "os"

func Detect() Kind {
	if k := detectFromEnv(os.Getenv("SHELL"), os.Getenv("ComSpec")); k != Unknown {
		return k
	}
	return Cmd
}
