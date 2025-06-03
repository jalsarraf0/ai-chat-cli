package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SendMessage(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		handler    http.HandlerFunc
		wantErr    bool
		wantResult string
	}{
		{
			name: "success",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				if _, err := w.Write([]byte(`{"reply":"pong"}`)); err != nil {
					t.Fatal(err)
				}
			},
			wantResult: "pong",
		},
		{
			name: "bad status",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			},
			wantErr: true,
		},
		{
			name: "invalid json",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				if _, err := w.Write([]byte("not-json")); err != nil {
					t.Fatal(err)
				}
			},
			wantErr: true,
		},
		{
			name: "empty reply",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				if _, err := w.Write([]byte(`{"reply":""}`)); err != nil {
					t.Fatal(err)
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			srv := httptest.NewServer(tt.handler)
			defer srv.Close()

			cli := NewClient(WithHTTPClient(srv.Client()), WithBaseURL(srv.URL))
			got, err := cli.SendMessage(context.Background(), "ping")
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected err %v, got %v", tt.wantErr, err)
			}
			if err == nil && got != tt.wantResult {
				t.Fatalf("expected %q, got %q", tt.wantResult, got)
			}
		})
	}
}

func ExampleClient() {
	cli := NewClient()
	_ = cli
	fmt.Println("ai response")
	// Output:
	// ai response
}

func TestNewClientOptions(t *testing.T) {
	t.Parallel()
	hc := &http.Client{}
	c := NewClient(WithHTTPClient(hc), WithBaseURL("x"), WithAPIKey("y"))
	if c.httpClient != hc || c.baseURL != "x" || c.apiKey != "y" {
		t.Fatalf("options not applied")
	}
}
