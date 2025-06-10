package cli_test

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	exe := filepath.Join(tmp, "ai-chat")
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	out, err := exec.Command("go", "build", "-o", exe, filepath.Join("..", "..", "..", "cmd", "ai-chat-cli")).CombinedOutput()
	if err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	return exe
}

func TestConfigCommandsIntegration(t *testing.T) {
	exe := buildBinary(t)
	cfg := filepath.Join(t.TempDir(), "c.yaml")

	cmd := exec.Command(exe, "--config", cfg, "config", "set", "openai_api_key", "abc")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("set: %v\n%s", err, out)
	}

	out, err := exec.Command(exe, "--config", cfg, "config", "get", "openai_api_key").Output()
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if string(out) == "" || string(out)[:3] != "abc" {
		t.Fatalf("unexpected get output: %s", out)
	}

	out, err = exec.Command(exe, "--config", cfg, "config", "list").Output()
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(out) == 0 {
		t.Fatalf("list empty")
	}

	out, err = exec.Command(exe, "--config", cfg, "config", "show").Output()
	if err != nil {
		t.Fatalf("show: %v", err)
	}
	if len(out) == 0 || !filepath.IsAbs(cfg) {
		t.Fatalf("show output bad: %s", out)
	}
}
