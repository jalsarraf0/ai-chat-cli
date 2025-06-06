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
