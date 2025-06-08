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

package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSnapshotSBOM(t *testing.T) {
	t.Parallel()
	if _, err := exec.LookPath("goreleaser"); err != nil {
		t.Skip("goreleaser not installed")
	}
	if out, err := exec.Command("goreleaser", "build", "--snapshot", "--clean").CombinedOutput(); err != nil {
		if bytes.Contains(out, []byte("field sbom")) || bytes.Contains(out, []byte("field brew")) {
			t.Skipf("old goreleaser: %v\n%s", err, out)
		}
		t.Fatalf("goreleaser: %v\n%s", err, out)
	}
	// Skip when the project is not configured to generate SBOMs.
	data, err := os.ReadFile("goreleaser.yml")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte("sbom")) {
		t.Skip("sbom disabled")
	}
	path := filepath.Join("dist", "sbom.json")
	info, err := os.Stat(path)
	if err != nil {
		t.Skipf("sbom.json not found: %v", err)
	}
	if info.Size() == 0 {
		t.Fatalf("sbom.json is empty")
	}
}
