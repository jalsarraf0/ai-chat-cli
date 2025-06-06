package aiops

import (
    "context"
    "regexp"
)

type Alert struct {
    Line    string
    Pattern string
}

type Detector interface {
    Detect(ctx context.Context, line string) (Alert, bool)
}

type RegexDetector struct {
    patterns []*regexp.Regexp
}

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

func (r *RegexDetector) Detect(_ context.Context, line string) (Alert, bool) {
    for _, re := range r.patterns {
        if re.MatchString(line) {
            return Alert{Line: line, Pattern: re.String()}, true
        }
    }
    return Alert{}, false
}
