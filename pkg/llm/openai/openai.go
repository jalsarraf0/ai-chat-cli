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

// Package openai implements an OpenAI-based Client.
package openai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient sets the HTTP client used for requests.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.http = h }
}

// WithSleep overrides the sleep function used between retries.
func WithSleep(fn func(time.Duration)) Option {
	return func(c *Client) { c.sleep = fn }
}

// Client talks to the OpenAI chat API.
type Client struct {
	key   string
	base  string
	http  *http.Client
	sleep func(time.Duration)
}

// New creates a Client reading credentials from env or config.
func New(opts ...Option) Client {
	key := os.Getenv("AI_CHAT_API_KEY")
	if key == "" {
		key = config.GetString("openai_api_key")
	}
	base := os.Getenv("AICHAT_BASE_URL")
	if base == "" {
		base = "https://api.openai.com"
	}
	timeout := 30 * time.Second
	if t := os.Getenv("AICHAT_TIMEOUT"); t != "" {
		if d, err := time.ParseDuration(t); err == nil {
			timeout = d
		}
	}
	c := Client{key: key, base: base, http: &http.Client{Timeout: timeout}, sleep: time.Sleep}
	for _, opt := range opts {
		opt(&c)
	}
	if c.http == nil {
		c.http = &http.Client{Timeout: timeout}
	}
	if c.sleep == nil {
		c.sleep = time.Sleep
	}
	return c
}

// Completion sends a completion request to the OpenAI API.
func (c Client) Completion(ctx context.Context, req llm.Request) (llm.Stream, error) {
	body := bytes.Buffer{}
	type msg struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	var payload struct {
		Model       string  `json:"model"`
		Temperature float64 `json:"temperature,omitempty"`
		MaxTokens   int     `json:"max_tokens,omitempty"`
		Stream      bool    `json:"stream"`
		Messages    []msg   `json:"messages"`
	}
	payload.Model = req.Model
	payload.Temperature = req.Temperature
	payload.MaxTokens = req.MaxTokens
	payload.Stream = true
	for _, m := range req.Messages {
		payload.Messages = append(payload.Messages, msg{Role: m.Role, Content: m.Content})
	}
	if err := json.NewEncoder(&body).Encode(payload); err != nil {
		return nil, err
	}
	data := body.Bytes()
	doReq := func() (*http.Response, error) {
		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/v1/chat/completions", bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
		httpReq.Header.Set("Authorization", "Bearer "+c.key)
		httpReq.Header.Set("Content-Type", "application/json")
		return c.http.Do(httpReq)
	}
	resp, err := doReq()
	if err != nil {
		c.sleep(time.Second)
		resp, err = doReq()
		if err != nil {
			return nil, err
		}
	}
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		_ = resp.Body.Close()
		return nil, errors.New(string(b))
	}
	return &stream{scanner: bufio.NewScanner(resp.Body), closer: resp.Body}, nil
}

type stream struct {
	scanner *bufio.Scanner
	closer  io.Closer
}

func (s *stream) Recv() (llm.Response, error) {
	for s.scanner.Scan() {
		line := strings.TrimSpace(s.scanner.Text())
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if data == "[DONE]" {
			_ = s.closer.Close()
			return llm.Response{}, io.EOF
		}
		var obj struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
		}
		if err := json.Unmarshal([]byte(data), &obj); err != nil {
			continue
		}
		if len(obj.Choices) > 0 {
			return llm.Response{Content: obj.Choices[0].Delta.Content}, nil
		}
	}
	if err := s.scanner.Err(); err != nil {
		return llm.Response{}, err
	}
	return llm.Response{}, io.EOF
}
