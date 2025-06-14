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
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

type asset struct {
	Name string
	Sum  string
}

func run(base string) error {
	var assets []asset
	patterns := []string{"templates/*.tmpl", "themes/*.json"}
	for _, p := range patterns {
		matches, _ := filepath.Glob(filepath.Join(base, p))
		for _, m := range matches {
			b, err := os.ReadFile(m) // #nosec G304 -- controlled glob
			if err != nil {
				return err
			}
			h := sha256.Sum256(b)
			name := filepath.ToSlash(m[len(base)+1:])
			assets = append(assets, asset{name, hex.EncodeToString(h[:])})
		}
	}
	sort.Slice(assets, func(i, j int) bool { return assets[i].Name < assets[j].Name })
	tmpl := template.Must(template.New("gen").Parse(`package assets

// Code generated by scripts/embedgen.go; DO NOT EDIT.

var assetSHA256 = map[string]string{
{{- range . }}
    "{{ .Name }}": "{{ .Sum }}",
{{- end }}
}
`))
	f, err := os.Create(filepath.Join(base, "assets_gen.go")) // #nosec G304 -- writing generated file
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	if err := tmpl.Execute(f, assets); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run("internal/assets"); err != nil {
		panic(err)
	}
}
