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

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestConfigCommands(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "key")
	config.Reset()
	cfg := newRootCmd()
	cfg.SetArgs([]string{"--config", dir + "/c.yaml", "config", "set", "openai_api_key", "abc"})
	if err := cfg.Execute(); err != nil {
		t.Fatalf("set: %v", err)
	}
	show := newRootCmd()
	out := new(bytes.Buffer)
	show.SetOut(out)
	show.SetArgs([]string{"--config", dir + "/c.yaml", "config", "show"})
	if err := show.Execute(); err != nil {
		t.Fatalf("show: %v", err)
	}
	outStr := out.String()
	want := dir + string(os.PathSeparator) + "c.yaml"
	want = filepath.ToSlash(want)
	if !strings.Contains(filepath.ToSlash(outStr), want) {
		t.Fatalf("missing path in output: %s", outStr)
	}
	edit := newRootCmd()
	edit.SetArgs([]string{"--config", dir + "/c.yaml", "config", "edit", "--dry-run"})
	if err := edit.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}

func TestConfigShowMissing(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "show"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestConfigSetInvalid(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "set", "model", "bad"})
	if err := c.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestConfigEditRun(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
	}
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	editor := filepath.Join(dir, "ed.sh")
	if err := os.WriteFile(editor, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("write editor: %v", err)
	}
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	t.Setenv("EDITOR", editor)
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}

func TestConfigEditDefaultEditor(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
	}
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	vi := filepath.Join(dir, "vi")
	if err := os.WriteFile(vi, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("write vi: %v", err)
	}
	t.Setenv("PATH", dir+":"+os.Getenv("PATH"))
	t.Setenv("EDITOR", "")
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}
