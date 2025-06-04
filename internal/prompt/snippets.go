package prompt

import _ "embed"

//go:embed dist/bash
var Bash string

//go:embed dist/zsh
var Zsh string

//go:embed dist/fish
var Fish string

//go:embed dist/powershell
var PowerShell string
