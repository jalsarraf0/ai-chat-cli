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

// Package openai wraps the go-openai client with a streaming helper.
package openai

import (
	"context"
	"io"
	"os"

	goopenai "github.com/sashabaranov/go-openai"
)

// Client streams completions from OpenAI.
type Client struct {
	api   *goopenai.Client
	model string
}

// New creates a Client using OPENAI_API_KEY and a default model.
func New() Client {
	key := os.Getenv("OPENAI_API_KEY")
	c := goopenai.NewClient(key)
	return Client{api: c, model: "gpt-3.5-turbo"}
}

// Stream sends a prompt and yields each token on a channel.
func (c Client) Stream(prompt string) (<-chan string, <-chan error) {
	out := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		req := goopenai.ChatCompletionRequest{
			Model:  c.model,
			Stream: true,
			Messages: []goopenai.ChatCompletionMessage{
				{Role: goopenai.ChatMessageRoleUser, Content: prompt},
			},
		}
		stream, err := c.api.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			errc <- err
			return
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				errc <- err
				return
			}
			if len(resp.Choices) > 0 {
				tok := resp.Choices[0].Delta.Content
				if tok != "" {
					out <- tok
				}
			}
		}
	}()
	return out, errc
}
