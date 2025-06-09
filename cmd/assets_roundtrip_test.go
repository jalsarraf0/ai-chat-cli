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
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAssetsRoundTrip(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	dest := filepath.Join(t.TempDir(), "out.json")
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err != nil {
		t.Fatalf("export: %v", err)
	}
	root = newRootCmd()
	out := new(bytes.Buffer)
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "themes/light.json"})
	if err := root.Execute(); err != nil {
		t.Fatalf("cat: %v", err)
	}
	b1, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !bytes.Equal(b1, out.Bytes()) {
		t.Fatalf("round trip mismatch")
	}
}
