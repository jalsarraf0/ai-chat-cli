package cmd

import (
    "bufio"
    "fmt"
    "os"

    "github.com/jalsarraf0/ai-chat-cli/internal/aiops"
    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"
)

type watchConfig struct {
    Patterns []string `yaml:"patterns"`
}

func newAIOpsCmd() *cobra.Command {
    cmd := &cobra.Command{Use: "aiops", Short: "AI-Ops helper"}
    cmd.AddCommand(newWatchCmd())
    return cmd
}

func newWatchCmd() *cobra.Command {
    var cfgPath string
    cmd := &cobra.Command{
        Use:   "watch --config <file>",
        Short: "Stream logs from stdin and detect anomalies",
        RunE: func(cmd *cobra.Command, _ []string) error {
            data, err := os.ReadFile(cfgPath)
            if err != nil {
                return err
            }
            var cfg watchConfig
            if err := yaml.Unmarshal(data, &cfg); err != nil {
                return err
            }
            det, err := aiops.NewRegexDetector(cfg.Patterns)
            if err != nil {
                return err
            }
            scanner := bufio.NewScanner(cmd.InOrStdin())
            for scanner.Scan() {
                alert, ok := det.Detect(cmd.Context(), scanner.Text())
                if ok {
                    fmt.Fprintf(cmd.OutOrStdout(), "%s: %s\n", alert.Pattern, alert.Line)
                }
            }
            return scanner.Err()
        },
    }
    cmd.Flags().StringVar(&cfgPath, "config", "", "config file")
    cmd.MarkFlagRequired("config")
    return cmd
}
