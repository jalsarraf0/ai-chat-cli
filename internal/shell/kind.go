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

// Kind enumerates supported shells.
type Kind string

// Supported shells.
const (
	// Unknown indicates the shell could not be detected.
	Unknown Kind = "unknown"
	// Bash shells.
	Bash Kind = "bash"
	// Zsh shell.
	Zsh Kind = "zsh"
	// Fish shell.
	Fish Kind = "fish"
	// PowerShell shell.
	PowerShell Kind = "powershell"
	// Cmd shell.
	Cmd Kind = "cmd"
)

func (k Kind) String() string { return string(k) }
