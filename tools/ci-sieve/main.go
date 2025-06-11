
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


//go:build tools

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var diffRange string

func init() {
	flag.StringVar(&diffRange, "range", "", "git diff range (e.g. base..HEAD)")
}

func main() {
	flag.Parse()
	files, err := changedFiles()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(files) == 0 {
		fmt.Println("no_changes")
		return
	}
	docsOnly := true
	for _, f := range files {
		if isDocPath(f) {
			continue
		}
		if !commentOnly(f) {
			docsOnly = false
			break
		}
	}
	if docsOnly {
		fmt.Println("docs_only")
		return
	}
	fmt.Println("code_change")
	os.Exit(1)
}

func gitDiffArgs() []string {
	if diffRange != "" {
		return []string{"diff", diffRange}
	}
	return []string{"diff", "--cached"}
}

func changedFiles() ([]string, error) {
	args := append(gitDiffArgs(), "--name-only")
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return nil, fmt.Errorf("git diff: %w", err)
	}
	lines := strings.Fields(string(out))
	return lines, nil
}

func isDocPath(p string) bool {
	if strings.HasPrefix(p, "docs/") {
		return true
	}
	switch filepath.Ext(p) {
	case ".md", ".rst", ".adoc", ".txt":
		return true
	}
	return false
}

var commentPrefixes = map[string][]string{
	".go":   {"//"},
	".sh":   {"#"},
	".bash": {"#"},
	".py":   {"#"},
	".rb":   {"#"},
	".yaml": {"#"},
	".yml":  {"#"},
	".toml": {"#"},
	".c":    {"//", "/*"},
	".cpp":  {"//", "/*"},
	".js":   {"//"},
	".ts":   {"//"},
}

var patchLine = regexp.MustCompile(`^\+(.*)`)

func commentOnly(path string) bool {
	args := gitDiffArgs()
	args = append(args, "-U0", path)
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "+++") || strings.HasPrefix(line, "---") || strings.HasPrefix(line, "diff") || strings.HasPrefix(line, "@@") {
			continue
		}
		m := patchLine.FindStringSubmatch(line)
		if len(m) != 2 {
			continue
		}
		content := strings.TrimSpace(m[1])
		if content == "" {
			continue
		}
		if isCommentLine(path, content) {
			continue
		}
		return false
	}
	return true
}

func isCommentLine(path, line string) bool {
	for _, p := range commentPrefixes[filepath.Ext(path)] {
		if strings.HasPrefix(line, p) {
			return true
		}
	}
	// generic comment characters
	return strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//")
}
