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
	"context"
	"runtime"
	"testing"
)

func TestRunCmd_LookPathStub(t *testing.T) {
	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(file string) (string, error) {
		if file == "echo" {
			if runtime.GOOS == "windows" {
				// cmd's internal 'echo' isn't an external binary,
				// so use cmd.exe and adapt the argument list below.
				return "C:\\Windows\\System32\\cmd.exe", nil
			}
			// Unix
			return "/bin/echo", nil
		}
		// Delegate to the original function for other binaries to
		// avoid interfering with parallel tests.
		return orig(file)
	}
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()

	var (
		stdout string
		stderr string
		err    error
	)

	if runtime.GOOS == "windows" {
		stdout, stderr, err = runCmd(context.Background(),
			"echo",
			[]string{"/C", "echo", "hi"},
		)
	} else {
		stdout, stderr, err = runCmd(context.Background(), "echo", []string{"hi"})
	}

	if err != nil {
		t.Fatalf("runCmd failed: %v", err)
	}

	want := "hi\n"
	if runtime.GOOS == "windows" {
		// cmd.exe adds trailing CRLF by default
		want = "hi\r\n"
	}
	if stdout != want || stderr != "" {
		t.Fatalf("unexpected output: %q / %q", stdout, stderr)
	}
}

func TestRunCmdLookPathError(t *testing.T) {
	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(string) (string, error) { return "", context.DeadlineExceeded }
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()
	if _, _, err := runCmd(context.Background(), "missing", nil); err == nil {
		t.Fatalf("expected error")
	}
}
