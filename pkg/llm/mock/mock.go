// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
}

// New creates a mock client that streams the given tokens.
func New(tokens ...string) Client {
	return Client{tokens: tokens}
}

// Completion returns a stream of predetermined tokens.
func (c Client) Completion(_ context.Context, _ llm.Request) (llm.Stream, error) {
	return &stream{tokens: c.tokens}, nil
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
