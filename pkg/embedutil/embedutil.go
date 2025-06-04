package embedutil

import (
	"bytes"
	"io/fs"
	"sort"

	"github.com/jalsarraf0/ai-chat-cli/internal/assets"
)

func List() []string {
	var names []string
	if err := fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, err error) error {
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

func Read(name string) ([]byte, error) { return assets.FS.ReadFile(name) }

func MustText(name string) string {
	data, err := Read(name)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}
