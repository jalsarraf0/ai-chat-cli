package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadSetEnv(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "config.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "envkey")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if GetString("openai_api_key") != "envkey" {
		t.Fatalf("env override failed")
	}
	if err := Set("model", "gpt-4"); err != nil {
		t.Fatalf("set: %v", err)
	}
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if len(b) == 0 {
		t.Fatalf("expected file content")
	}
}

func TestValidation(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := Set("model", "bad"); err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestPathAndSave(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if Path() != file {
		t.Fatalf("path mismatch")
	}
	if err := Save(); err != nil {
		t.Fatalf("save: %v", err)
	}
}

func TestDefaultPathXDG(t *testing.T) {
	Reset()
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	t.Setenv("XDG_CONFIG_HOME", dir)
	p := defaultPath()
	want := filepath.Join(dir, "ai-chat", "config.yaml")
	if p != want {
		t.Fatalf("want %s got %s", want, p)
	}
}

func TestDefaultPathHome(t *testing.T) {
	Reset()
	t.Setenv("XDG_CONFIG_HOME", "")
	p := defaultPath()
	if !strings.Contains(p, ".config/ai-chat/config.yaml") {
		t.Fatalf("unexpected %s", p)
	}
}

func TestDefaultPathFallback(t *testing.T) {
	Reset()
	t.Setenv("XDG_CONFIG_HOME", "")
	oldHome := os.Getenv("HOME")
	t.Setenv("HOME", "")
	p := defaultPath()
	if !strings.Contains(p, "ai-chat/config.yaml") {
		t.Fatalf("unexpected %s", p)
	}
	if oldHome != "" {
		t.Setenv("HOME", oldHome)
	}
}

func TestSetValidModel(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := Set("model", "gpt-3.5-turbo"); err != nil {
		t.Fatalf("valid model: %v", err)
	}
}

func TestLoadExistingFile(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := Set("model", "gpt-4"); err != nil {
		t.Fatalf("set: %v", err)
	}
	Reset()
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("reload: %v", err)
	}
	if GetString("model") != "gpt-4" {
		t.Fatalf("want gpt-4")
	}
}

func TestLoadMissingKey(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	if err := Load(file); err == nil {
		t.Fatalf("expected error")
	}
}

func TestSaveError(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	path = filepath.Join(dir, "no", "c.yaml")
	if err := Save(); err == nil {
		t.Fatalf("expected error")
	}
}
