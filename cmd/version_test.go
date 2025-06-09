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
	"fmt"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"short", []string{"--short"}, "1.2.3\n"},
		{"full", nil, "1.2.3 abc now\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", "/bin/bash")
			cmd := newVersionCmd("1.2.3", "abc", "now")
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetArgs(tt.args)
			if err := cmd.Execute(); err != nil {
				t.Fatalf("execute: %v", err)
			}
			if buf.String() != tt.want {
				t.Fatalf("want %q got %q", tt.want, buf.String())
			}
		})
	}
}

type versionErrWriter struct{}

func (versionErrWriter) Write(_ []byte) (int, error) { return 0, fmt.Errorf("werr") }

func TestVersionWriteError(t *testing.T) {
	cmd := newVersionCmd("1.2.3", "abc", "now")
	cmd.SetOut(versionErrWriter{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
