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

// Package llm defines language model interfaces.
package llm

import "context"

// Message represents a chat message.
type Message struct {
	Role    string
	Content string
}

// Request describes a completion query.
type Request struct {
	Model       string
	Temperature float64
	MaxTokens   int
	Messages    []Message
}

// Response is a single chunk of completion text.
type Response struct {
	Content string
}

// Stream yields completion tokens.
type Stream interface {
	Recv() (Response, error)
}

// Client sends completion requests to a language model.
type Client interface {
	Completion(ctx context.Context, req Request) (Stream, error)
	// ListModels returns all available model identifiers.
	ListModels(ctx context.Context) ([]string, error)
}
