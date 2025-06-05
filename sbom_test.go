package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
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
		t.Fatalf("goreleaser: %v\n%s", err, out)
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
