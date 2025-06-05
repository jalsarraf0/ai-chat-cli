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
