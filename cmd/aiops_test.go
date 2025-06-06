package cmd

import (
    "bytes"
    "os"
    "testing"
)

func TestWatchCmd(t *testing.T) {
    t.Parallel()
    cfg := []byte("patterns:\n  - ERROR\n")
    f, err := os.CreateTemp(t.TempDir(), "cfg.yaml")
    if err != nil { t.Fatal(err) }
    if _, err := f.Write(cfg); err != nil { t.Fatal(err) }
    if err := f.Close(); err != nil { t.Fatal(err) }

    logFile, err := os.Open("../testdata/logs/anomaly.log")
    if err != nil { t.Fatal(err) }
    out := new(bytes.Buffer)
    cmd := newWatchCmd()
    cmd.SetArgs([]string{"--config", f.Name()})
    cmd.SetIn(logFile)
    cmd.SetOut(out)
    if err := cmd.Execute(); err != nil { t.Fatalf("run: %v", err) }
    if out.String() == "" { t.Fatalf("expected alert") }
}
