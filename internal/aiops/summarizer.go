// Copyright 2025 AI Chat
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
