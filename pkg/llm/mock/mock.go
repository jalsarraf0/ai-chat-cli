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

// Package mock provides test implementations of llm.Client.
package mock

import (
	"context"
	"io"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

// Client returns predetermined tokens.
type Client struct {
	tokens []string
	models []string
}

// New creates a mock client that streams the given tokens.
func New(tokens ...string) Client {
	return Client{tokens: tokens, models: []string{"gpt-4.1-nano", "gpt-3.5-turbo"}}
}

// Completion returns a stream of predetermined tokens.
func (c Client) Completion(_ context.Context, _ llm.Request) (llm.Stream, error) {
	return &stream{tokens: c.tokens}, nil
}

// ListModels returns a fixed set of model names.
func (c Client) ListModels(context.Context) ([]string, error) {
	return c.models, nil
}

type stream struct {
	tokens []string
	i      int
}

func (s *stream) Recv() (llm.Response, error) {
	if s.i >= len(s.tokens) {
		return llm.Response{}, io.EOF
	}
	t := s.tokens[s.i]
	s.i++
	return llm.Response{Content: t}, nil
}
