// Copyright 2025 The ai-chat-cli Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build integration

package openai

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func TestLiveCompletion(t *testing.T) {
	if os.Getenv("AI_CHAT_API_KEY") == "" {
		t.Skip("AI_CHAT_API_KEY not set")
	}
	c := New()
	s, err := c.Completion(context.Background(), llm.Request{
		Model:    "gpt-3.5-turbo",
		Messages: []llm.Message{{Role: "user", Content: "ping"}},
	})
	if err != nil {
		t.Fatalf("request: %v", err)
	}
	_, err = s.Recv()
	if err != nil && err != io.EOF {
		t.Fatalf("recv: %v", err)
	}
}
