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

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jalsarraf0/ai-chat-cli/internal/openai"
)

// Controller streams tokens from the LLM and sends messages to the program.
type Controller struct {
	client openai.Client
	prog   *tea.Program
}

// NewController creates a controller bound to the program.
func NewController(c openai.Client, p *tea.Program) Controller {
	return Controller{client: c, prog: p}
}

// Stream spawns a goroutine to stream a prompt and forward tokens.
func (c Controller) Stream(prompt string) {
	out, _ := c.client.Stream(prompt)
	go func() {
		for tok := range out {
			c.prog.Send(llmTokenMsg{Token: tok})
		}
		c.prog.Send(llmTokenMsg{Done: true})
	}()
}
