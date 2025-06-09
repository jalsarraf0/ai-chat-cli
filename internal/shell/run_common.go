// Copyright 2025 AI Chat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
