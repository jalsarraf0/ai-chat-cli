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

// Package theme provides colour palettes.
package theme

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
)

// Palette represents a colour palette.
type Palette struct {
	Background string `json:"background"`
}

// Load reads the palette with optional name. When name is empty it selects a
// default based on COLORTERM ("light" selects the light palette).
func Load(name string) Palette {
	if name == "" {
		ct := strings.ToLower(os.Getenv("COLORTERM"))
		if strings.Contains(ct, "light") {
			name = "themes/light.json"
		} else {
			name = "themes/dark.json"
		}
	}
	data, err := embedutil.Read(name)
	if err != nil {
		data = []byte(`{"background":""}`)
	}
	var p Palette
	_ = json.Unmarshal(data, &p)
	return p
}
