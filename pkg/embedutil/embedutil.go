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

// Package embedutil exposes helpers for embedded assets.
package embedutil

import (
	"bytes"
	"io/fs"
	"sort"

	"github.com/jalsarraf0/ai-chat-cli/internal/assets"
)

var (
	assetsFS fs.FS = assets.FS
	walkDir        = fs.WalkDir
)

// List returns the names of all embedded files.
func List() []string {
	var names []string
	if err := walkDir(assetsFS, ".", func(path string, d fs.DirEntry, _ error) error {
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
