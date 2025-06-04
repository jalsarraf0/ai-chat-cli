//go:build windows

package config

import (
	"os"
	"path/filepath"
)

func defaultPathImpl() string {
	return filepath.Join(os.Getenv("APPDATA"), "ai-chat", "config.yaml")
}
