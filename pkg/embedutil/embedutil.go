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

// Package embedutil exposes helpers for embedded assets.
package embedutil

import (
	"bytes"
	"io/fs"
	"sort"

	"github.com/jalsarraf0/ai-chat-cli/internal/assets"
)

// List returns the names of all embedded files.
func List() []string {
	var names []string
	if err := fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, _ error) error {
		if !d.IsDir() {
			names = append(names, path)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	sort.Strings(names)
	return names
}

// Read returns a file's bytes from the embedded filesystem.
func Read(name string) ([]byte, error) { return assets.FS.ReadFile(name) }

// MustText returns a file's contents as a string and panics on error.
func MustText(name string) string {
	data, err := Read(name)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}
