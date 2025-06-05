package testhttp

import "net/http"

// RoundTripFunc allows custom HTTP behavior in tests.
type RoundTripFunc func(*http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
