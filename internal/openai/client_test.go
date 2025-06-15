package openai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gpt "github.com/sashabaranov/go-openai"
)

func TestStream(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Fatalf("path=%s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		fmt.Fprintf(w, "data: {\"choices\":[{\"delta\":{\"content\":\"hi\"}}]}\n\n")
		fmt.Fprintf(w, "data: {\"choices\":[{\"delta\":{\"content\":\" there\"}}]}\n\n")
		fmt.Fprintf(w, "data: [DONE]\n\n")
	}))
	defer srv.Close()

	old := newClient
	newClient = func(key string) *gpt.Client {
		cfg := gpt.DefaultConfig(key)
		cfg.BaseURL = srv.URL + "/v1"
		return gpt.NewClientWithConfig(cfg)
	}
	defer func() { newClient = old }()

	t.Setenv("OPENAI_API_KEY", "k")
	ch := Stream("hi")
	var out string
	for tok := range ch {
		out += tok
	}
	if out != "hi there" {
		t.Fatalf("got %q", out)
	}
}

func TestStreamNoKey(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	ch := Stream("hi")
	if _, ok := <-ch; ok {
		t.Fatalf("expected closed channel")
	}
}
