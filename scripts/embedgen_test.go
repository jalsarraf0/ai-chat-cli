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

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func copyDir(dst, src string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		ds := filepath.Join(dst, e.Name())
		ss := filepath.Join(src, e.Name())
		if e.IsDir() {
			if err := os.MkdirAll(ds, 0o755); err != nil {
				return err
			}
			if err := copyDir(ds, ss); err != nil {
				return err
			}
		} else {
			b, err := os.ReadFile(ss)
			if err != nil {
				return err
			}
			if err := os.WriteFile(ds, b, 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}

func TestRunSuccess(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(cwd) }()
	if err := run("internal/assets"); err != nil {
		t.Fatal(err)
	}
}

func TestMainFunction(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(cwd) }()
	main()
}

func TestMainFunctionError(t *testing.T) {
	tmp := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(cwd) }()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	main()
}

func TestRunMissingDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := run(filepath.Join(tmpDir, "internal/assets"))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestRunReadError(t *testing.T) {
	tmpDir := t.TempDir()
	// copy assets to tmpDir
	src := filepath.Join("..", "internal", "assets")
	dst := filepath.Join(tmpDir, "internal", "assets")
	if err := os.MkdirAll(dst, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := copyDir(dst, src); err != nil {
		t.Fatal(err)
	}
	badFile := filepath.Join(dst, "templates", "default.tmpl")
	if err := os.Remove(badFile); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(badFile, 0o755); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Remove(badFile)
		// restore original file
		orig := filepath.Join(src, "templates", "default.tmpl")
		b, _ := os.ReadFile(orig)
		_ = os.WriteFile(badFile, b, 0o644)
	}()
	if err := run(dst); err == nil {
		t.Fatal("expected error")
	}
}

func TestRunCreateError(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "internal", "assets")
	if err := os.MkdirAll(filepath.Join(tmpDir, "internal"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filePath, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := run(filePath); err == nil {
		t.Fatal("expected error")
	}
}
