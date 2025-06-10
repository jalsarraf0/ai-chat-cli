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
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
)

func TestAssetsCommands(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	root := newRootCmd()
	out := new(bytes.Buffer)
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "list"})
	if err := root.Execute(); err != nil {
		t.Fatalf("list: %v", err)
	}
	lines := bytes.Split(bytes.TrimSpace(out.Bytes()), []byte("\n"))
	if len(lines) != 4 {
		t.Fatalf("want 4 lines got %d", len(lines))
	}

	// cat
	root = newRootCmd()
	out.Reset()
	root.SetOut(out)
	root.SetArgs([]string{"--config", cfg, "assets", "cat", "templates/system.tmpl"})
	if err := root.Execute(); err != nil {
		t.Fatalf("cat: %v", err)
	}
	want, _ := embedutil.Read("templates/system.tmpl")
	if !bytes.Equal(out.Bytes(), want) {
		t.Fatalf("cat output mismatch")
	}

	// export
	dest := filepath.Join(t.TempDir(), "x", "y.json")
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err != nil {
		t.Fatalf("export: %v", err)
	}
	b1, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("read dest: %v", err)
	}
	b2, _ := embedutil.Read("themes/light.json")
	if !bytes.Equal(b1, b2) {
		t.Fatalf("export mismatch")
	}
	// overwrite with --force
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest, "--force"})
	if err := root.Execute(); err != nil {
		t.Fatalf("export force: %v", err)
	}
}

func TestAssetsExportError(t *testing.T) {
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	dest := filepath.Join(t.TempDir(), "dup.json")
	if err := os.WriteFile(dest, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
