//go:build unit

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
		io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n")
		io.WriteString(w, "data: [DONE]\n\n")
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "event: noop\n\n")
		io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"x\"}}]}\n\n")
		io.WriteString(w, "data: [DONE]\n\n")
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "data: {bad}\n\n")
		io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n")
		io.WriteString(w, "data: [DONE]\n\n")
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "data: {\"choices\":[{\"delta\":{\"content\":\"ok\"}}]}\n\n")
		io.WriteString(w, "data: [DONE]\n\n")
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
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "data: {")
		conn, _, _ := w.(http.Hijacker).Hijack()
		conn.Close()
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
	rt := testhttp.RoundTripFunc(func(r *http.Request) (*http.Response, error) {
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
	os.Setenv("AI_CHAT_API_KEY", "tok")
	os.Setenv("AICHAT_BASE_URL", "")
	os.Setenv("AICHAT_TIMEOUT", "")
	t.Cleanup(func() {
		os.Setenv("AI_CHAT_API_KEY", oldKey)
		os.Setenv("AICHAT_BASE_URL", oldBase)
		os.Setenv("AICHAT_TIMEOUT", oldTimeout)
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
	os.Setenv("AI_CHAT_API_KEY", "tok")
	os.Setenv("AICHAT_TIMEOUT", "5s")
	t.Cleanup(func() {
		os.Setenv("AI_CHAT_API_KEY", oldKey)
		os.Setenv("AICHAT_TIMEOUT", oldTimeout)
	})
	c := New()
	if c.http.Timeout != 5*time.Second {
		t.Fatalf("timeout %v", c.http.Timeout)
	}
}

func TestNewNilOptions(t *testing.T) {
	oldKey := os.Getenv("AI_CHAT_API_KEY")
	os.Setenv("AI_CHAT_API_KEY", "tok")
	t.Cleanup(func() { os.Setenv("AI_CHAT_API_KEY", oldKey) })
	c := New(WithHTTPClient(nil), WithSleep(nil))
	if c.http == nil || c.sleep == nil {
		t.Fatalf("nil not defaulted")
	}
}
