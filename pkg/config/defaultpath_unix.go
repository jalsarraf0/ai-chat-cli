//go:build !windows

package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/adrg/xdg"
)

func defaultPathImpl() string {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return filepath.Join(dir, "ai-chat", "config.yaml")
	}
	if u, err := user.Current(); err == nil {
		return filepath.Join(u.HomeDir, ".config", "ai-chat", "config.yaml")
	}
	return filepath.Join(xdg.Home, ".config", "ai-chat", "config.yaml")
}
