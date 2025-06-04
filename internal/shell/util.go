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
