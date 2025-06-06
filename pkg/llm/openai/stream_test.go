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

//go:build unit

package openai

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func TestCompletionStream(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			t.Fatalf("missing auth")
		}
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"hi\"}}]}\n\n")
		io.WriteString(w, "data: [DONE]\n\n")
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := New()
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "user", Content: "q"}}})
	if err != nil {
		t.Fatalf("req: %v", err)
	}
	r1, err := s.Recv()
	if err != nil || r1.Content != "hi" {
		t.Fatalf("token: %v %v", r1, err)
	}
	_, err = s.Recv()
	if err != io.EOF {
		t.Fatalf("want EOF")
	}
}

func TestCompletionHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad")
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := New()
	_, err := c.Completion(context.Background(), llm.Request{Model: "m"})
	if err == nil || !strings.Contains(err.Error(), "bad") {
		t.Fatalf("expected error, got %v", err)
	}
}
