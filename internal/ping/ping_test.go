package ping

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	t.Parallel()
	if err := Ping(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
