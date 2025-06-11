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
	"strings"
	"testing"
)

func TestAssetsExportErrors(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "k")
	cfg := filepath.Join(t.TempDir(), "c.yaml")
	dest := filepath.Join(t.TempDir(), "out.json")

	root := newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "bogus", dest})
	if err := root.Execute(); err == nil {
		t.Fatalf("expected error")
	}

	// existing file without --force
	if err := os.WriteFile(dest, []byte("x"), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
	root = newRootCmd()
	root.SetArgs([]string{"--config", cfg, "assets", "export", "themes/light.json", dest})
	err := root.Execute()
	if err == nil || !strings.Contains(err.Error(), "exists") {
		t.Fatalf("want exists error, got %v", err)
	}
}
