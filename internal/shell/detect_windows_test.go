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

import "testing"

func TestDetectWindows(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("ComSpec", `C:\\Windows\\System32\\cmd.exe`)
	if got := Detect(); got != Cmd {
		t.Fatalf("want cmd got %v", got)
	}
	t.Setenv("SHELL", `C:\\Program Files\\PowerShell\\pwsh.exe`)
	if got := Detect(); got != PowerShell {
		t.Fatalf("want pwsh got %v", got)
	}
}
