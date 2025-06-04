//go:build unit

package testhttp

import (
	"net/http"
	"testing"
)

func TestRoundTripFunc(t *testing.T) {
	rt := RoundTripFunc(func(r *http.Request) (*http.Response, error) {
		return nil, http.ErrUseLastResponse
	})
	_, err := rt.RoundTrip(&http.Request{})
	if err != http.ErrUseLastResponse {
		t.Fatalf("got %v", err)
	}
}
