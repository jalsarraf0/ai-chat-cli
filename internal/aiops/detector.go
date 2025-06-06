// Copyright 2025 The ai-chat-cli Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
