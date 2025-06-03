package shell

import (
    "bytes"
    "context"
    "os/exec"
)

var lookPath = exec.LookPath

func runCmd(ctx context.Context, name string, args []string) (string, string, error) {
    cmd := exec.CommandContext(ctx, name, args...)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return stdout.String(), stderr.String(), err
}
