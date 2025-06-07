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
	"sort"
	"strings"
)

// Summarizer condenses a text into a short summary.
type Summarizer interface {
	Summarize(text string) string
}

// TFIDFSummarizer is a naive word frequency summarizer.
type TFIDFSummarizer struct{}

// NewTFIDFSummarizer returns a new TFIDFSummarizer.
func NewTFIDFSummarizer() *TFIDFSummarizer { return &TFIDFSummarizer{} }

// Summarize returns the most frequent words from the given text.
func (s *TFIDFSummarizer) Summarize(text string) string {
	words := strings.Fields(strings.ToLower(text))
	freq := map[string]int{}
	for _, w := range words {
		freq[w]++
	}
	type pair struct {
		word  string
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for w, c := range freq {
		pairs = append(pairs, pair{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })
	n := 5
	if len(pairs) < n {
		n = len(pairs)
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = pairs[i].word
	}
	return strings.Join(out, " ")
}
