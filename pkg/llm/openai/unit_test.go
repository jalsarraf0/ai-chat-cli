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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jalsarraf0/ai-chat-cli/internal/testhttp"
	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func newUnitClient(srv *httptest.Server, sleep func(time.Duration)) Client {
	return New(WithHTTPClient(srv.Client()), WithSleep(sleep))
}

func TestEnvOverride(t *testing.T) {
	var auth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
	}))
	defer srv.Close()

	t.Setenv("AI_CHAT_API_KEY", "token")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "user", Content: "q"}}})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "ok" {
		t.Fatalf("recv: %v %v", r, err)
	}
	if auth != "Bearer token" {
		t.Fatalf("auth %q", auth)
	}
}

func TestHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "fail", http.StatusBadRequest)
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	_, err := c.Completion(context.Background(), llm.Request{Model: "m"})
	if err == nil || !strings.Contains(err.Error(), "fail") {
		t.Fatalf("want error, got %v", err)
	}
}

func TestIgnoreNonDataLines(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := io.WriteString(w, "event: noop\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"x\"}}]}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "u", Content: "q"}}})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "x" {
		t.Fatalf("recv: %v %v", r, err)
	}
}

func TestInvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := io.WriteString(w, "data: {bad}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "u", Content: "q"}}})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "ok" {
		t.Fatalf("recv: %v %v", r, err)
	}
}

func TestRetry(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
		if _, err := io.WriteString(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("write: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	first := true
	slept := false
	rt := testhttp.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
		if first {
			first = false
			return nil, io.ErrUnexpectedEOF
		}
		return srv.Client().Transport.RoundTrip(r)
	})
	client := &http.Client{Transport: rt}
	c := New(WithHTTPClient(client), WithSleep(func(time.Duration) { slept = true }))
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "u", Content: "q"}}})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "ok" {
		t.Fatalf("recv: %v %v", r, err)
	}
	if !slept {
		t.Fatalf("sleep not called")
	}
}

func TestNewFromConfig(t *testing.T) {
	config.Reset()
	dir := t.TempDir()
	file := filepath.Join(dir, "c.yaml")
	t.Setenv("AICHAT_OPENAI_API_KEY", "k")
	if err := config.Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := config.Set("openai_api_key", "cfg"); err != nil {
		t.Fatalf("set: %v", err)
	}
	t.Setenv("AI_CHAT_API_KEY", "")
	c := New()
	if c.key != "cfg" {
		t.Fatalf("want cfg got %s", c.key)
	}
}

func TestRecvScannerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := io.WriteString(w, "data: {"); err != nil {
			t.Fatalf("write: %v", err)
		}
		conn, _, _ := w.(http.Hijacker).Hijack()
		if err := conn.Close(); err != nil {
			t.Fatalf("close: %v", err)
		}
	}))
	defer srv.Close()
	t.Setenv("AI_CHAT_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	s, err := c.Completion(context.Background(), llm.Request{Model: "m", Messages: []llm.Message{{Role: "u", Content: "q"}}})
	if err != nil {
		t.Fatalf("completion: %v", err)
	}
	_, err = s.Recv()
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestRetryFails(t *testing.T) {
	rt := testhttp.RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
		return nil, io.ErrUnexpectedEOF
	})
	client := &http.Client{Transport: rt}
	t.Setenv("AI_CHAT_API_KEY", "k")
	c := New(WithHTTPClient(client), WithSleep(func(time.Duration) {}))
	_, err := c.Completion(context.Background(), llm.Request{Model: "m"})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestNewDefaults(t *testing.T) {
	config.Reset()
	oldKey := os.Getenv("AI_CHAT_API_KEY")
	oldBase := os.Getenv("AICHAT_BASE_URL")
	oldTimeout := os.Getenv("AICHAT_TIMEOUT")
	if err := os.Setenv("AI_CHAT_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_BASE_URL", ""); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_TIMEOUT", ""); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("AI_CHAT_API_KEY", oldKey)
		_ = os.Setenv("AICHAT_BASE_URL", oldBase)
		_ = os.Setenv("AICHAT_TIMEOUT", oldTimeout)
	})
	c := New()
	if c.base != "https://api.openai.com" {
		t.Fatalf("base %s", c.base)
	}
	if c.http == nil || c.http.Timeout != 30*time.Second {
		t.Fatalf("timeout %v", c.http)
	}
	if c.sleep == nil {
		t.Fatalf("sleep nil")
	}
}

func TestNewTimeoutEnv(t *testing.T) {
	oldKey := os.Getenv("AI_CHAT_API_KEY")
	oldTimeout := os.Getenv("AICHAT_TIMEOUT")
	if err := os.Setenv("AI_CHAT_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_TIMEOUT", "5s"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("AI_CHAT_API_KEY", oldKey)
		_ = os.Setenv("AICHAT_TIMEOUT", oldTimeout)
	})
	c := New()
	if c.http.Timeout != 5*time.Second {
		t.Fatalf("timeout %v", c.http.Timeout)
	}
}

func TestNewNilOptions(t *testing.T) {
	oldKey := os.Getenv("AI_CHAT_API_KEY")
	if err := os.Setenv("AI_CHAT_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() { _ = os.Setenv("AI_CHAT_API_KEY", oldKey) })
	c := New(WithHTTPClient(nil), WithSleep(nil))
	if c.http == nil || c.sleep == nil {
		t.Fatalf("nil not defaulted")
	}
}
