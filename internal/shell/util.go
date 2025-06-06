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

package shell

import "path/filepath"

func fromBaseName(p string) string {
	return filepath.Base(p)
}

func classify(base string) Kind {
	switch base {
	case "bash", "bash.exe":
		return Bash
	case "zsh", "zsh.exe":
		return Zsh
	case "fish", "fish.exe":
		return Fish
	case "pwsh", "pwsh.exe", "powershell", "powershell.exe":
		return PowerShell
	case "cmd", "cmd.exe":
		return Cmd
	default:
		return Unknown
	}
}
