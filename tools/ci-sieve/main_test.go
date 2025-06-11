//go:build tools

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIsDocPath(t *testing.T) {
	cases := map[string]bool{
		"docs/readme.md": true,
		"README.md":      true,
		"file.txt":       true,
		"src/main.go":    false,
	}
	for f, want := range cases {
		if got := isDocPath(f); got != want {
			t.Fatalf("isDocPath(%s)=%v want %v", f, got, want)
		}
	}
}

func TestCommentOnly(t *testing.T) {
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v: %v\n%s", args, err, out)
		}
	}
	run("git", "init")
	run("git", "config", "user.email", "you@example.com")
	run("git", "config", "user.name", "Your Name")
	path := filepath.Join(dir, "main.go")
	os.WriteFile(path, []byte("package main\n"), 0o644)
	run("git", "add", "main.go")
	run("git", "commit", "-m", "init")

	// comment only change
	os.WriteFile(path, []byte("package main\n// hi"), 0o644)
	run("git", "add", "main.go")
	cwd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(cwd)
	if !commentOnly("main.go") {
		t.Fatalf("expected true")
	}
	// code change
	os.WriteFile(path, []byte("package main\nvar x = 1"), 0o644)
	run("git", "add", "main.go")
	if commentOnly("main.go") {
		t.Fatalf("expected false")
	}
}

func TestIsCommentLine(t *testing.T) {
	if !isCommentLine("main.go", "// foo") {
		t.Fatalf("expected true")
	}
	if isCommentLine("main.go", "var x int") {
		t.Fatalf("expected false")
	}
}

func TestMainDocsOnly(t *testing.T) {
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v: %v\n%s", args, err, out)
		}
	}
	run("git", "init")
	run("git", "config", "user.email", "you@example.com")
	run("git", "config", "user.name", "Your Name")
	path := filepath.Join(dir, "README.md")
	os.WriteFile(path, []byte("hi"), 0o644)
	run("git", "add", "README.md")
	run("git", "commit", "-m", "init")

	os.WriteFile(path, []byte("hi there"), 0o644)
	run("git", "add", "README.md")
	cwd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(cwd)
	if docsOnly := mainReturn(); !docsOnly {
		t.Fatalf("expected docs-only")
	}

	// code change should flip result
	code := filepath.Join(dir, "file.go")
	os.WriteFile(code, []byte("package main"), 0o644)
	run("git", "add", "file.go")
	run("git", "commit", "-m", "code")
	os.WriteFile(code, []byte("package main\nvar x int"), 0o644)
	run("git", "add", "file.go")
	if mainReturn() {
		t.Fatalf("expected code-change")
	}
}

func TestChangedFiles(t *testing.T) {
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v: %v\n%s", args, err, out)
		}
	}
	run("git", "init")
	run("git", "config", "user.email", "you@example.com")
	run("git", "config", "user.name", "Your Name")
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("hi"), 0o644)
	run("git", "add", "a.txt")
	run("git", "commit", "-m", "init")
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("bye"), 0o644)
	run("git", "add", "a.txt")
	cwd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(cwd)
	files, err := changedFiles()
	if err != nil || len(files) != 1 || files[0] != "a.txt" {
		t.Fatalf("changedFiles unexpected: %v %v", files, err)
	}
}

func mainReturn() bool {
	files, err := changedFiles()
	if err != nil || len(files) == 0 {
		return false
	}
	for _, f := range files {
		if !isDocPath(f) && !commentOnly(f) {
			return false
		}
	}
	return true
}
