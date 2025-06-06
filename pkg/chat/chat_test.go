// Copyright 2025 The ai-chat-cli Authors
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

package chat

import (
	"context"
	"testing"
)

type errorClient struct{}

func (errorClient) Ping(_ context.Context) error              { return context.Canceled }
func (errorClient) Version(_ context.Context) (string, error) { return "", context.Canceled }

func TestClientImplementations(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		client      Client
		wantPing    bool
		wantVersion string
	}{
		{"mock", mockClient{}, false, "mock"},
		{"error", errorClient{}, true, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.client.Ping(context.Background())
			if tt.wantPing && err == nil {
				t.Fatalf("expected ping error")
			}
			if !tt.wantPing && err != nil {
				t.Fatalf("ping err %v", err)
			}
			v, err := tt.client.Version(context.Background())
			if tt.wantPing && err == nil {
				t.Fatalf("expected version error")
			}
			if !tt.wantPing && (err != nil || v != tt.wantVersion) {
				t.Fatalf("got %q err %v", v, err)
			}
		})
	}
}
func TestNewMockClient(t *testing.T) {
	t.Parallel()
	c := NewMockClient()
	if v, err := c.Version(context.Background()); err != nil || v != "mock" {
		t.Fatalf("NewMockClient unexpected: v=%s err=%v", v, err)
	}
}
