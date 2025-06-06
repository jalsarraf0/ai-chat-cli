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

package embedutil

import (
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()
	got := List()
	want := []string{
		"templates/default.tmpl",
		"templates/system.tmpl",
		"themes/dark.json",
		"themes/light.json",
	}
	if len(got) != len(want) {
		t.Fatalf("want %d got %d", len(want), len(got))
	}
	for i, n := range want {
		if got[i] != n {
			t.Fatalf("want %s got %s", n, got[i])
		}
	}
}

func TestMustText(t *testing.T) {
	t.Parallel()
	txt := MustText("templates/default.tmpl")
	if !strings.Contains(txt, "{{.User}}") {
		t.Fatalf("missing variable in %q", txt)
	}
}

func TestMustTextPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	MustText("missing.txt")
}
