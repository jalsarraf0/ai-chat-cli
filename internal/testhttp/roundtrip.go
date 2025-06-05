// Package testhttp provides HTTP helpers for tests.
package testhttp

import "net/http"

// RoundTripFunc allows custom HTTP behavior in tests.
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip executes the wrapped function.
func (f RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
