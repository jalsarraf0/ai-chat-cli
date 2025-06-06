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

//go:build windows

package shell

import (
	"context"
)

// Run executes a command via the user's shell on Windows.
func Run(ctx context.Context, cmd string, args ...string) (string, string, error) {
	lookPathMu.RLock()
	lp := lookPath
	lookPathMu.RUnlock()
	pwsh, _ := lp("pwsh")
	if pwsh == "" {
		pwsh, _ = lp("powershell")
	}
	if pwsh != "" {
		return runCmd(ctx, pwsh, append([]string{"-NoProfile", "-Command", cmd}, args...))
	}
	return runCmd(ctx, "cmd", append([]string{"/C", cmd}, args...))
}
