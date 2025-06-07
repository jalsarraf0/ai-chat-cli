// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

	// #nosec G204 -- path is resolved via lookPath; not user-supplied.
	cmd := exec.CommandContext(ctx, path, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	return stdout.String(), stderr.String(), err
}
