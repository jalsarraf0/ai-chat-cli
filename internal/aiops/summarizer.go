package aiops

import (
    "sort"
    "strings"
)

type Summarizer interface {
    Summarize(text string) string
}

type TFIDFSummarizer struct{}

func NewTFIDFSummarizer() *TFIDFSummarizer { return &TFIDFSummarizer{} }

func (s *TFIDFSummarizer) Summarize(text string) string {
    words := strings.Fields(strings.ToLower(text))
    freq := map[string]int{}
    for _, w := range words {
        freq[w]++
    }
    type pair struct{ word string; count int }
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
