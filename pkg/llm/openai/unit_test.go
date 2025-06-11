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

package openai

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
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

	t.Setenv("OPENAI_API_KEY", "token")
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
	t.Setenv("OPENAI_API_KEY", "k")
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
	t.Setenv("OPENAI_API_KEY", "k")
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
	t.Setenv("OPENAI_API_KEY", "k")
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
	t.Setenv("OPENAI_API_KEY", "k")
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
	t.Setenv("OPENAI_API_KEY", "k")
	if err := config.Load(file); err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := config.Set("openai_api_key", "cfg"); err != nil {
		t.Fatalf("set: %v", err)
	}
	t.Setenv("OPENAI_API_KEY", "")
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
	t.Setenv("OPENAI_API_KEY", "k")
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
	t.Setenv("OPENAI_API_KEY", "k")
	c := New(WithHTTPClient(client), WithSleep(func(time.Duration) {}))
	_, err := c.Completion(context.Background(), llm.Request{Model: "m"})
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestNewDefaults(t *testing.T) {
	config.Reset()
	oldKey := os.Getenv("OPENAI_API_KEY")
	oldBase := os.Getenv("AICHAT_BASE_URL")
	oldTimeout := os.Getenv("AICHAT_TIMEOUT")
	if err := os.Setenv("OPENAI_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_BASE_URL", ""); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_TIMEOUT", ""); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("OPENAI_API_KEY", oldKey)
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
	oldKey := os.Getenv("OPENAI_API_KEY")
	oldTimeout := os.Getenv("AICHAT_TIMEOUT")
	if err := os.Setenv("OPENAI_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	if err := os.Setenv("AICHAT_TIMEOUT", "5s"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("OPENAI_API_KEY", oldKey)
		_ = os.Setenv("AICHAT_TIMEOUT", oldTimeout)
	})
	c := New()
	if c.http.Timeout != 5*time.Second {
		t.Fatalf("timeout %v", c.http.Timeout)
	}
}

func TestNewNilOptions(t *testing.T) {
	oldKey := os.Getenv("OPENAI_API_KEY")
	if err := os.Setenv("OPENAI_API_KEY", "tok"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() { _ = os.Setenv("OPENAI_API_KEY", oldKey) })
	c := New(WithHTTPClient(nil), WithSleep(nil))
	if c.http == nil || c.sleep == nil {
		t.Fatalf("nil not defaulted")
	}
}

func TestListModels(t *testing.T) {

	t.Setenv("OPENAI_API_KEY", "k")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"data": [{"id": "gpt-a"}, {"id": "gpt-b"}]}`)
	}))
	defer srv.Close()
	t.Setenv("OPENAI_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})

	models, err := c.ListModels(context.Background())
	if err != nil {
		t.Fatalf("list: %v", err)
	}

	if len(models) == 0 {
		t.Fatalf("no models returned")
	}
	if sort.StringsAreSorted(models) == false {
		t.Fatalf("models not sorted")
	}
	found := false
	for _, m := range models {
		if m == "gpt-4.1-nano" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("gpt-4.1-nano missing: %v", models)
		if len(models) != 2 || models[0] != "gpt-a" || models[1] != "gpt-b" {
			t.Fatalf("models %v", models)
		}
	}
}
func TestListModelsHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "bad", http.StatusBadRequest)
	}))
	defer srv.Close()
	t.Setenv("OPENAI_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	_, err := c.ListModels(context.Background())
	if err == nil || !strings.Contains(err.Error(), "bad") {
		t.Fatalf("want error got %v", err)

	}
}

func TestListModelsDecodeError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "not-json")
	}))
	defer srv.Close()
	t.Setenv("OPENAI_API_KEY", "k")
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	c := newUnitClient(srv, func(time.Duration) {})
	if _, err := c.ListModels(context.Background()); err == nil {
		t.Fatalf("expected decode error")
	}
}

func TestListModelsNoKey(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		called = true
	}))
	defer srv.Close()
	t.Setenv("AICHAT_BASE_URL", srv.URL)
	t.Setenv("OPENAI_API_KEY", "")
	c := newUnitClient(srv, func(time.Duration) {})
	models, err := c.ListModels(context.Background())
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if called {
		t.Fatalf("unexpected http call")
	}
	if len(models) == 0 {
		t.Fatalf("no models")
	}
	found := false
	for _, m := range models {
		if m == "gpt-4.1-nano" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("gpt-4.1-nano missing")
	}
}




