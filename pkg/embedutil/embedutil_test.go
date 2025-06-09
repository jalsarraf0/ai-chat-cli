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

func TestReadError(t *testing.T) {
	t.Parallel()
	if _, err := Read("does-not-exist"); err == nil {
		t.Fatalf("expected error")
	}
}
