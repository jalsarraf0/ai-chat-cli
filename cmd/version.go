// Package cmd provides CLI commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func newVersionCmd(version, commit, date string) *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, _ []string) {
			if short {
				if _, err := fmt.Fprintln(cmd.OutOrStdout(), version); err != nil {
					cmd.Println("Error:", err)
				}
				return
			}
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s %s %s\n", version, commit, date); err != nil {
				cmd.Println("Error:", err)
			}
		},
	}
	cmd.Flags().BoolVarP(&short, "short", "s", false, "short output")
	return cmd
}
