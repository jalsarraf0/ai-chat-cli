// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package aiops provides utilities for anomaly detection and summarisation of
// log streams.
package aiops

import (
	"context"
	"regexp"
)

// Alert represents a single log line that matched an anomaly pattern.
type Alert struct {
	Line    string
	Pattern string
}

// Detector looks for anomalies in a log line.
type Detector interface {
	Detect(ctx context.Context, line string) (Alert, bool)
}

// RegexDetector implements Detector using regular expression patterns.
type RegexDetector struct {
	patterns []*regexp.Regexp
}

// NewRegexDetector returns a RegexDetector that matches any of the provided
// patterns.
func NewRegexDetector(patterns []string) (*RegexDetector, error) {
	rs := make([]*regexp.Regexp, len(patterns))
	for i, p := range patterns {
		r, err := regexp.Compile(p)
		if err != nil {
			return nil, err
		}
		rs[i] = r
	}
	return &RegexDetector{patterns: rs}, nil
}

// Detect checks a log line and returns an Alert if any pattern matches.
func (r *RegexDetector) Detect(_ context.Context, line string) (Alert, bool) {
	for _, re := range r.patterns {
		if re.MatchString(line) {
			return Alert{Line: line, Pattern: re.String()}, true
		}
	}
	return Alert{}, false
}
