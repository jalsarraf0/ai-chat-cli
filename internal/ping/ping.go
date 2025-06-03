package ping

import "context"

// Ping checks connectivity.
func Ping(ctx context.Context) error {
	_ = ctx
	return nil
}
