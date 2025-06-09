// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestSnapshotSBOM(t *testing.T) {
	t.Parallel()
	if _, err := exec.LookPath("goreleaser"); err != nil {
		t.Skip("goreleaser not installed")
	}
	out, err := exec.Command("goreleaser", "build", "--snapshot", "--clean").CombinedOutput()
	if err != nil {
		if strings.Contains(err.Error(), "exec format error") {
			t.Skipf("goreleaser wrong arch: %v\n%s", err, out)
		}
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
	matches, err := filepath.Glob("dist/*.tar.gz")
	if err != nil || len(matches) == 0 {
		t.Fatalf("archive not found: %v", err)
	}
	f, err := os.Open(matches[0])
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = f.Close()
	}()
	gz, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	tr := tar.NewReader(gz)
	found := false
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if hdr.Name == "sbom.json" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("sbom.json not found in %s", matches[0])
	}
}
