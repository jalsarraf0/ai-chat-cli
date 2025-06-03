package shell

import (
	"bytes"
	"context"
	"os/exec"
)

// Run executes the provided command string in the detected shell and returns
// collected stdout and stderr bytes.
func Run(ctx context.Context, cmd string) ([]byte, []byte, error) {
	name, path, err := Detect()
	if err != nil {
		return nil, nil, err
	}

	var c *exec.Cmd
	switch name {
	case "bash", "zsh", "sh", "fish":
		c = exec.CommandContext(ctx, path, "-c", cmd)
	case "cmd":
		c = exec.CommandContext(ctx, path, "/C", cmd)
	case "powershell":
		c = exec.CommandContext(ctx, path, "-NoProfile", "-NonInteractive", "-Command", cmd)
	default:
		return nil, nil, ErrUnsupported
	}

	var outBuf, errBuf bytes.Buffer
	c.Stdout = &outBuf
	c.Stderr = &errBuf

	if err := c.Run(); err != nil {
		return outBuf.Bytes(), errBuf.Bytes(), err
	}
	return outBuf.Bytes(), errBuf.Bytes(), nil
}
