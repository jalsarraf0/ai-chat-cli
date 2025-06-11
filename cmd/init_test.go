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
)

func TestInitRunsScript(t *testing.T) {
	script := filepath.Join(t.TempDir(), "setup.sh")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho ok"), 0o755); err != nil {
		t.Fatalf("write: %v", err)
	}
	setupScript = script
	cmd := newInitCmd()
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if out.String() != "ok\n" {
		t.Fatalf("out=%q", out.String())
	}
}

func TestInitRelativeScript(t *testing.T) {
	exe, err := os.Executable()
	if err != nil {
		t.Skip(err)
	}
	dir := filepath.Dir(exe)
	script := filepath.Join(dir, "setup.sh")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho hi"), 0o755); err != nil {
		t.Fatalf("write: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(script) })
	setupScript = "setup.sh"
	cmd := newInitCmd()
	out := new(bytes.Buffer)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("exec: %v", err)
	}
	if out.String() != "hi\n" {
		t.Fatalf("out=%q", out.String())
	}
}
