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

package cli_test

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLISmoke(t *testing.T) {
	exe := buildBinary(t)
	cfg := filepath.Join(t.TempDir(), "c.yaml")

	if out, err := exec.Command(exe, "--help").CombinedOutput(); err != nil {
		t.Fatalf("--help: %v\n%s", err, out)
	}

	if out, err := exec.Command(exe, "--config", cfg, "config", "list").CombinedOutput(); err != nil {
		t.Fatalf("list: %v\n%s", err, out)
	} else if len(out) != 0 {
		t.Logf("list output: %s", out)
	}

	if out, err := exec.Command(exe, "--config", cfg, "config", "set", "openai_api_key", "dummy").CombinedOutput(); err != nil {
		t.Fatalf("set api key: %v\n%s", err, out)
	}

	if out, err := exec.Command(exe, "--config", cfg, "config", "set", "foo", "bar").CombinedOutput(); err != nil {
		t.Fatalf("set foo: %v\n%s", err, out)
	}

	out, err := exec.Command(exe, "--config", cfg, "config", "get", "foo").Output()
	if err != nil {
		t.Fatalf("get foo: %v", err)
	}
	if strings.TrimSpace(string(out)) != "bar" {
		t.Fatalf("unexpected get output: %s", out)
	}
}
