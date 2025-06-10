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
	"context"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/chat"
)

type stubClient struct{}

type errClient struct{}

type errWriter struct{}

func (errWriter) Write(_ []byte) (int, error) { return 0, context.DeadlineExceeded }

func (stubClient) Ping(_ context.Context) error              { return nil }
func (stubClient) Version(_ context.Context) (string, error) { return "", nil }
func (errClient) Ping(_ context.Context) error               { return context.DeadlineExceeded }
func (errClient) Version(_ context.Context) (string, error)  { return "", context.DeadlineExceeded }

func TestPingCommand(t *testing.T) {
	tests := []struct {
		name    string
		shell   string
		verbose bool
		client  chat.Client
		wantErr bool
	}{
		{"normal", "/bin/bash", false, stubClient{}, false},
		{"verbose", "/bin/zsh", true, stubClient{}, false},
		{"error", "/bin/bash", false, errClient{}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("SHELL", tt.shell)
			verbose = tt.verbose
			defer func() { verbose = false }()
			cmd := newPingCmd(tt.client)
			out := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)
			cmd.SetOut(out)
			cmd.SetErr(errBuf)
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if !tt.wantErr && out.String() != "pong\n" {
				t.Fatalf("out=%q", out.String())
			}
			if tt.verbose && errBuf.String() == "" {
				t.Fatalf("expected shell info")
			}
		})
	}
}

func TestPingWriteError(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	cmd := newPingCmd(stubClient{})
	cmd.SetOut(errWriter{})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error")
	}
}
