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

func TestFromBaseName(t *testing.T) {
	cases := map[string]string{
		"/bin/bash":          "bash",
		"C:/Windows/cmd.exe": "cmd.exe",
	}
	for in, want := range cases {
		if got := fromBaseName(in); got != want {
			t.Fatalf("%s => %s want %s", in, got, want)
		}
	}
}
