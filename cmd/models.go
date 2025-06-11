package cmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
	"github.com/spf13/cobra"
)

func newModelsCmd(c llm.Client) *cobra.Command {
	var list bool
	cmd := &cobra.Command{
		Use:   "models",
		Short: "List available models",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !list {
				list = true
			}
			if list {
				models, err := c.ListModels(context.Background())
				if err != nil {
					return err
				}
				sort.Strings(models)
				for _, m := range models {
					if _, err := fmt.Fprintln(cmd.OutOrStdout(), m); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&list, "list", false, "list models")
	return cmd
}
