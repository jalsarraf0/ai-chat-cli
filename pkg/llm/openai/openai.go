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

// Package openai implements an OpenAI client.
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
	"sort"
	"strings"
	"time"

	"github.com/jalsarraf0/ai-chat-cli/pkg/config"
	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

var defaultModels = []string{
	"gpt-4o",
	"gpt-4o-mini",
	"gpt-4o-audio-preview",
	"gpt-4o-2024-05-13",
	"gpt-4.1",
	"gpt-4.1-mini",
	"gpt-4.1-nano",
	"gpt-4.1-2025-04-14",
	"gpt-4",
	"gpt-4-32k",
	"gpt-4-turbo",
	"gpt-4-turbo-preview",
	"gpt-4-vision-preview",
	"gpt-4-0314",
	"gpt-4-0613",
	"gpt-4-0125-preview",
	"gpt-3.5-turbo",
	"gpt-3.5-turbo-16k",
	"gpt-3.5-turbo-0125",
	"gpt-3.5-turbo-1106",
	"text-embedding-3-large",
	"text-embedding-3-small",
	"text-embedding-ada-002",
	"whisper-1",
	"dall-e-3",
	"moderation-latest",
	"moderation-v1",
	"gpt-4o-nano",
	"gpt-image-1",
}

func sorted(m map[string]struct{}) []string {
	models := make([]string, 0, len(m))
	for id := range m {
		models = append(models, id)
	}
	sort.Strings(models)
	return models
}

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
	key := os.Getenv("OPENAI_API_KEY")
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

// ListModels retrieves available model identifiers from the OpenAI API.
func (c Client) ListModels(ctx context.Context) ([]string, error) {
	uniq := map[string]struct{}{}
	for _, m := range defaultModels {
		uniq[m] = struct{}{}
	}

	if c.key == "" {
		return sorted(uniq), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/v1/models", nil)
	if err != nil {
		return sorted(uniq), err
	}
	req.Header.Set("Authorization", "Bearer "+c.key)
	resp, err := c.http.Do(req)
	if err != nil {
		return sorted(uniq), err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return sorted(uniq), errors.New(string(b))
	}
	var data struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return sorted(uniq), err
	}

	for _, m := range data.Data {
		uniq[m.ID] = struct{}{}

	uniq := map[string]struct{}{}
	for _, m := range defaultModels {
		uniq[m] = struct{}{}
	}
	for _, m := range data.Data {
		uniq[m.ID] = struct{}{}
	}

	models := make([]string, 0, len(uniq))
	for id := range uniq {
		models = append(models, id)

	}

	return sorted(uniq), nil

}
