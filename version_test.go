package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestStampedVersion(t *testing.T) {
	t.Parallel()
	tmp := t.TempDir()
	exe := filepath.Join(tmp, "ai-chat")
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	ldflags := "-X github.com/jalsarraf0/ai-chat-cli/cmd.Version=1.0.0 -X github.com/jalsarraf0/ai-chat-cli/cmd.Commit=abc123 -X github.com/jalsarraf0/ai-chat-cli/cmd.Date=now"
	if out, err := exec.Command("go", "build", "-o", exe, "-ldflags", ldflags, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	cmd := exec.Command(exe, "version")
	cmd.Env = append(cmd.Env, "AICHAT_OPENAI_API_KEY=dummy")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("run: %v", err)
	}
	if !strings.Contains(string(out), "1.0.0 abc123 now") {
		t.Fatalf("unexpected output: %s", out)
	}
}
