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
