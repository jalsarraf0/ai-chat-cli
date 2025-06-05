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
