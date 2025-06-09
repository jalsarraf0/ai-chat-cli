// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
