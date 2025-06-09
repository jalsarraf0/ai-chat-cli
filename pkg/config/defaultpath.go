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

package config

import (
	"os"
	"path/filepath"
)

func defaultPathImpl() string {
	if p := os.Getenv("AI_CHAT_CONFIG"); p != "" {
		return p
	}
	base := os.Getenv("XDG_CONFIG_HOME")
	if base == "" {
		if h := os.Getenv("HOME"); h != "" {
			base = filepath.Join(h, ".config")
		} else {
			base = os.TempDir()
		}
	}
	return filepath.Join(base, "ai-chat", "ai-chat.yaml")
}
