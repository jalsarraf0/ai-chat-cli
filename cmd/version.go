/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
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
                Run: func(cmd *cobra.Command, args []string) {
                        if short {
                                fmt.Fprintln(cmd.OutOrStdout(), version)
                                return
                        }
                        fmt.Fprintf(cmd.OutOrStdout(), "%s %s %s\n", version, commit, date)
                },
        }
        cmd.Flags().BoolVarP(&short, "short", "s", false, "short output")
        return cmd
}

