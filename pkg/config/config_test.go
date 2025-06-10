// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
       want := filepath.Join(dir, "ai-chat-cli", "config.yaml")
	if p != want {
		t.Fatalf("want %s got %s", want, p)
	}
}

func TestDefaultPathHome(t *testing.T) {
	Reset()
	t.Setenv("XDG_CONFIG_HOME", "")
	p := defaultPath()
       if !strings.Contains(p, ".config/ai-chat-cli/config.yaml") {
		t.Fatalf("unexpected %s", p)
	}
}

func TestDefaultPathFallback(t *testing.T) {
	Reset()
	t.Setenv("XDG_CONFIG_HOME", "")
	oldHome := os.Getenv("HOME")
	t.Setenv("HOME", "")
	p := defaultPath()
       if !strings.Contains(p, "ai-chat-cli/config.yaml") {
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

func TestLoadInvalidYAML(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	if err := os.WriteFile(file, []byte(": bad"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := Load(file); err == nil {
		t.Fatalf("expected error")
	}
}

func TestLoadDirError(t *testing.T) {
	Reset()
	dir := t.TempDir()
	path := filepath.Join(dir, "c.yaml")
	if err := os.Mkdir(path, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	if err := Load(path); err == nil {
		t.Fatalf("expected error")
	}
}

func TestGetters(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	v.Set("x", 1.5)
	if GetFloat64("x") != 1.5 {
		t.Fatalf("float")
	}
	if GetInt("missing") != 0 {
		t.Fatalf("int default")
	}
}

func TestAllAndIsSet(t *testing.T) {
	Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	if err := Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := Set("foo", "bar"); err != nil {
		t.Fatalf("set: %v", err)
	}
	if !IsSet("foo") {
		t.Fatalf("expected key")
	}
	if Get("foo") != "bar" {
		t.Fatalf("get value")
	}
	all := All()
	if all["foo"] != "bar" {
		t.Fatalf("all map")
	}
}
