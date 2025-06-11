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

package mock

import (
	"context"
	"io"
	"testing"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

func TestStream(t *testing.T) {
	t.Parallel()
	c := New("hi", "there")
	s, err := c.Completion(context.Background(), llm.Request{})
	if err != nil {
		t.Fatalf("completion: %v", err)
	}
	r, err := s.Recv()
	if err != nil || r.Content != "hi" {
		t.Fatalf("first token: %v %v", r, err)
	}
	r, err = s.Recv()
	if err != nil || r.Content != "there" {
		t.Fatalf("second token: %v %v", r, err)
	}
	_, err = s.Recv()
	if err != io.EOF {
		t.Fatalf("want EOF")
	}
}

func TestListModels(t *testing.T) {
	c := New()
	models, err := c.ListModels(context.Background())
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(models) == 0 {
		t.Fatalf("no models")
	}
}
