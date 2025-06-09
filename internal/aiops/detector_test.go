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
	"context"
	"testing"
)

func TestRegexDetector(t *testing.T) {
	t.Parallel()
	det, err := NewRegexDetector([]string{"ERROR", "FATAL"})
	if err != nil {
		t.Fatalf("new: %v", err)
	}
	tests := []struct {
		line string
		ok   bool
	}{
		{"info: ok", false},
		{"ERROR: bad", true},
		{"FATAL: crash", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.line, func(t *testing.T) {
			t.Parallel()
			_, got := det.Detect(context.Background(), tt.line)
			if got != tt.ok {
				t.Fatalf("Detect=%v want %v", got, tt.ok)
			}
		})
	}
}
