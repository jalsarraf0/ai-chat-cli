package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/jalsarraf0/ai-chat-cli/internal/prompt"
	"github.com/jalsarraf0/ai-chat-cli/internal/shell"
	"github.com/spf13/cobra"
)

func newInitCmd(root *cobra.Command) *cobra.Command {
	var dryRun bool
	var noPrompt bool
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Install completion and prompt snippet",
		RunE: func(cmd *cobra.Command, args []string) error {
			sh := shell.Detect()
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			dir := filepath.Join(home, ".ai-chat")
			ext := map[shell.Kind]string{
				shell.Bash:       "bash",
				shell.Zsh:        "zsh",
				shell.Fish:       "fish",
				shell.PowerShell: "ps1",
			}[sh]
			if ext == "" {
				return fmt.Errorf("unsupported shell %s", sh)
			}
			comp := filepath.Join(dir, "completion."+ext)
			if !dryRun {
				if err := os.MkdirAll(dir, 0o755); err != nil {
					return err
				}
				r := NewRootCmd()
				r.SetArgs([]string{"completion", sh.String(), "--out", comp})
				r.SetOut(io.Discard)
				r.SetErr(io.Discard)
				if err := r.Execute(); err != nil {
					return err
				}
			}
			rcFile := rcPath(home, sh)
			snippet := sourceLine(sh, comp)
			if !noPrompt {
				snippet += promptFor(sh)
			}
			if dryRun {
				_, err = fmt.Fprintln(cmd.OutOrStdout(), snippet)
			} else {
				f, err := openAppend(rcFile, 0o644)
				if err != nil {
					return err
				}
				defer f.Close()
				if _, err := f.WriteString(snippet); err != nil {
					return err
				}
			}
			_, err = fmt.Fprintf(cmd.OutOrStdout(), "✅ Completions installed for %s\n", sh)
			return err
		},
	}
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "print actions only")
	cmd.Flags().BoolVar(&noPrompt, "no-prompt", false, "skip prompt snippet")
	return cmd
}

func rcPath(home string, sh shell.Kind) string {
	switch sh {
	case shell.Bash:
		return filepath.Join(home, ".bashrc")
	case shell.Zsh:
		return filepath.Join(home, ".zshrc")
	case shell.Fish:
		return filepath.Join(home, ".config", "fish", "config.fish")
	case shell.PowerShell:
		return filepath.Join(home, "Documents", "WindowsPowerShell", "profile.ps1")
	default:
		return ""
	}
}

func sourceLine(sh shell.Kind, file string) string {
	if sh == shell.PowerShell {
		return fmt.Sprintf(". \"%s\"\n", file)
	}
	return fmt.Sprintf("source %s\n", file)
}

func promptFor(sh shell.Kind) string {
	switch sh {
	case shell.Bash:
		return "\n" + prompt.Bash + "\n"
	case shell.Zsh:
		return "\n" + prompt.Zsh + "\n"
	case shell.Fish:
		return "\n" + prompt.Fish + "\n"
	case shell.PowerShell:
		return "\n" + prompt.PowerShell + "\n"
	default:
		return ""
	}
}

func openAppend(path string, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, perm)
}
