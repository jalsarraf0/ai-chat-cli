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
	"testing"
)

func TestWatchCmd(t *testing.T) {
	t.Parallel()
	cfg := []byte("patterns:\n  - ERROR\n")
	f, err := os.CreateTemp(t.TempDir(), "cfg.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write(cfg); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	logFile, err := os.Open("../test/integration/logs/anomaly.log")
	if err != nil {
		t.Fatal(err)
	}
	out := new(bytes.Buffer)
	cmd := newWatchCmd()
	cmd.SetArgs([]string{"--config", f.Name()})
	cmd.SetIn(logFile)
	cmd.SetOut(out)
	if err := cmd.Execute(); err != nil {
		t.Fatalf("run: %v", err)
	}
	if out.String() == "" {
		t.Fatalf("expected alert")
	}
}

func TestWatchCmdError(t *testing.T) {
	t.Parallel()
	cmd := newWatchCmd()
	cmd.SetArgs([]string{"--config", "nope.yaml"})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
