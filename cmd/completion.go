package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newCompletionCmd(root *cobra.Command) *cobra.Command {
	var out string
	cmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion",
		Args:  cobra.ExactArgs(1),
		Example: "  ai-chat completion bash --out /tmp/ai.bash\n" +
			"  ai-chat completion zsh",
		RunE: func(cmd *cobra.Command, args []string) error {
			sh := args[0]
			if out == "" {
				out = filepath.Join("dist", "completion", sh)
			}
			if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
				return err
			}
			f, err := os.Create(out)
			if err != nil {
				return err
			}
			defer func() {
				if cerr := f.Close(); cerr != nil {
					cmd.PrintErrf("close file: %v\n", cerr)
				}
			}()
			if err := generateCompletion(root, sh, f); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&out, "out", "", "output file")
	return cmd
}

func generateCompletion(root *cobra.Command, sh string, w io.Writer) error {
	switch sh {
	case "bash":
		return root.GenBashCompletion(w)
	case "zsh":
		return root.GenZshCompletion(w)
	case "fish":
		return root.GenFishCompletion(w, true)
	case "powershell":
		return root.GenPowerShellCompletionWithDesc(w)
	default:
		return fmt.Errorf("unsupported shell: %s", sh)
	}
}
