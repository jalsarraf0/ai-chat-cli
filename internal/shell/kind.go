package shell

// Kind enumerates supported shells.
type Kind string

const (
    Unknown     Kind = "unknown"
    Bash        Kind = "bash"
    Zsh         Kind = "zsh"
    Fish        Kind = "fish"
    PowerShell  Kind = "powershell"
    Cmd         Kind = "cmd"
)

func (k Kind) String() string { return string(k) }
