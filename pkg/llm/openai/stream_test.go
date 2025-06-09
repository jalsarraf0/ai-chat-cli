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
		if _, err := io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"hi\"}}]}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := io.WriteString(w, "bad"); err != nil {
			t.Fatalf("write: %v", err)
		}
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
