// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including but not limited to the rights to
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

package theme

import (
        "testing"
)

func TestLoadDefault(t *testing.T) {
    t.Setenv("COLORTERM", "light")
	p := Load("")
	if p.Background != "#ffffff" {
		t.Fatalf("want light palette")
	}
}

func TestLoadNamed(t *testing.T) {
	p := Load("themes/dark.json")
	if p.Background != "#000000" {
		t.Fatalf("want dark palette")
	}
}

func TestLoadDefaultDark(t *testing.T) {
    t.Setenv("COLORTERM", "dumb")
	p := Load("")
	if p.Background != "#000000" {
		t.Fatalf("want dark palette")
	}
}

func TestLoadMissing(t *testing.T) {
	p := Load("not/a/file.json")
	if p.Background != "" {
		t.Fatalf("want empty palette")
	}
}
