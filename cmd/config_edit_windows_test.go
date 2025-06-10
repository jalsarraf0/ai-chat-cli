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
	"os"
	"path/filepath"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
)

func TestConfigEditRunWindows(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	config.Reset()
	editor := filepath.Join(dir, "ed.bat")
	// simple batch file that exits successfully
	if err := os.WriteFile(editor, []byte("@echo off\r\n"), 0o755); err != nil {
		t.Fatalf("write editor: %v", err)
	}
	c := newRootCmd()
	c.SetArgs([]string{"--config", filepath.Join(dir, "c.yaml"), "config", "edit"})
	t.Setenv("EDITOR", editor)
	if err := c.Execute(); err != nil {
		t.Fatalf("edit: %v", err)
	}
}
