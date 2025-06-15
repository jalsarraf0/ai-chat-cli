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

package openai

import (
	"context"
	"os"

	gpt "github.com/sashabaranov/go-openai"
)

var newClient = gpt.NewClient

// Stream returns a channel streaming tokens from the OpenAI API using the
// default key from the environment. It closes the channel when the stream
// finishes or on error.
func Stream(prompt string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		key := os.Getenv("OPENAI_API_KEY")
		if key == "" {
			return
		}
		client := newClient(key)
		req := gpt.ChatCompletionRequest{
			Model:  gpt.GPT3Dot5Turbo,
			Stream: true,
			Messages: []gpt.ChatCompletionMessage{
				{Role: gpt.ChatMessageRoleUser, Content: prompt},
			},
		}
		stream, err := client.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			return
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				return
			}
			if len(resp.Choices) == 0 {
				continue
			}
			ch <- resp.Choices[0].Delta.Content
		}
	}()
	return ch
}
