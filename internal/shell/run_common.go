package shell

import (
	"bytes"
	"context"
	"os/exec"
	"sync"
)

// lookPath is a test-hookable alias for exec.LookPath.
var (
	lookPath   = exec.LookPath
	lookPathMu sync.RWMutex
)

// runCmd executes a command respecting context and returns stdout/stderr strings.
func runCmd(ctx context.Context, name string, args []string) (string, string, error) {
	// Resolve the absolute binary path via lookPath (allows test stubbing).
	lookPathMu.RLock()
	lp := lookPath
	lookPathMu.RUnlock()
	path, err := lp(name)
	if err != nil {
		return "", "", err
	}

	cmd := exec.CommandContext(ctx, path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	return stdout.String(), stderr.String(), err
}
