package chat

import (
    "context"
    "testing"
)

type errorClient struct{}

func (errorClient) Ping(ctx context.Context) error { return context.Canceled }
func (errorClient) Version(ctx context.Context) (string, error) { return "", context.Canceled }

func TestMockClient(t *testing.T) {
    t.Parallel()
    c := mockClient{}
    if err := c.Ping(context.Background()); err != nil {
        t.Fatalf("ping err: %v", err)
    }
    v, err := c.Version(context.Background())
    if err != nil || v != "mock" {
        t.Fatalf("version got %q err %v", v, err)
    }
}

func TestErrorClient(t *testing.T) {
    t.Parallel()
    c := errorClient{}
    if err := c.Ping(context.Background()); err == nil {
        t.Fatalf("expected error")
    }
    if _, err := c.Version(context.Background()); err == nil {
        t.Fatalf("expected error")
    }
}
