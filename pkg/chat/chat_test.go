package chat

import (
	"context"
	"testing"
)

type errorClient struct{}

func (errorClient) Ping(ctx context.Context) error              { return context.Canceled }
func (errorClient) Version(ctx context.Context) (string, error) { return "", context.Canceled }

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
	c := NewMockClient()
	if err := c.Ping(context.Background()); err != nil {
		t.Fatalf("ping: %v", err)
	}
	if v, err := c.Version(context.Background()); err != nil || v != "mock" {
		t.Fatalf("version: %v %q", err, v)
	}
}
