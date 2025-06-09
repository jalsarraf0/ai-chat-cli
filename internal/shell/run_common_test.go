// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
