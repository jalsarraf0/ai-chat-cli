// Copyright 2025 AI Chat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
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
	"os"
	"path/filepath"
	"testing"
)

func TestAssetsListWriteError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "list"})
	root.SetOut(failWriter{})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestAssetsCatError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "missing.txt"})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestAssetsExportMkdirError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	base := filepath.Join(t.TempDir(), "dir")
	if err := os.MkdirAll(base, 0o755); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(base, "file")
	if err := os.WriteFile(file, []byte{}, 0o600); err != nil {
		t.Fatal(err)
	}
	dest := filepath.Join(file, "out.json")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
