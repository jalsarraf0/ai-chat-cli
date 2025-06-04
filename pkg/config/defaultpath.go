package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func defaultPathImpl() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "ai-chat", "config.yaml")
	}
	if runtime.GOOS == "windows" {
		if app := os.Getenv("APPDATA"); app != "" {
			return filepath.Join(app, "ai-chat", "config.yaml")
		}
	}
	if home := os.Getenv("HOME"); home != "" {
		return filepath.Join(home, ".config", "ai-chat", "config.yaml")
	}
	return filepath.Join(os.TempDir(), "ai-chat", "config.yaml")
}
