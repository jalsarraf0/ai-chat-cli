package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "config", Short: "Manage configuration"}
	cmd.AddCommand(newConfigShowCmd())
	cmd.AddCommand(newConfigSetCmd())
	cmd.AddCommand(newConfigEditCmd())
	return cmd
}

func newConfigShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := os.ReadFile(configPath())
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), string(b))
			return err
		},
	}
}

func newConfigSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return config.Set(args[0], args[1])
		},
	}
}

func newConfigEditCmd() *cobra.Command {
	var dryRun bool
	c := &cobra.Command{
		Use:   "edit",
		Short: "Edit configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRun {
				return nil
			}
			editor := os.Getenv("EDITOR")
			if editor == "" {
				if runtime.GOOS == "windows" {
					editor = "notepad"
				} else {
					editor = "vi"
				}
			}
			e := exec.Command(editor, configPath())
			e.Stdin = os.Stdin
			e.Stdout = os.Stdout
			e.Stderr = os.Stderr
			return e.Run()
		},
	}
	c.Flags().BoolVar(&dryRun, "dry-run", false, "skip launching editor")
	return c
}

func configPath() string {
	return config.Path()
}
