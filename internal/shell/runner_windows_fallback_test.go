// Copyright 2025 The ai-chat-cli Authors
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

//go:build windows

package shell

import (
	"context"
	"os/exec"
	"testing"
)

// TestRunWindowsFallback verifies cmd is used when PowerShell is absent.
func TestRunWindowsFallback(t *testing.T) {
	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(name string) (string, error) {
		if name == "cmd" {
			return exec.LookPath("cmd")
		}
		return "", exec.ErrNotFound
	}
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()
	if _, _, err := Run(context.Background(), "echo", "hi"); err != nil {
		t.Fatalf("run: %v", err)
	}
}
