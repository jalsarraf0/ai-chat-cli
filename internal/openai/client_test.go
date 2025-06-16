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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gpt "github.com/sashabaranov/go-openai"
)

func TestStream(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Fatalf("path=%s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		if _, err := fmt.Fprintf(w, "data: {\"choices\":[{\"delta\":{\"content\":\"hi\"}}]}\n\n"); err != nil {
			t.Fatalf("%v", err)
		}
		if _, err := fmt.Fprintf(w, "data: {\"choices\":[{\"delta\":{\"content\":\" there\"}}]}\n\n"); err != nil {
			t.Fatalf("%v", err)
		}
		if _, err := fmt.Fprintf(w, "data: [DONE]\n\n"); err != nil {
			t.Fatalf("%v", err)
		}
	}))
	defer srv.Close()

	old := newClient
	newClient = func(key string) *gpt.Client {
		cfg := gpt.DefaultConfig(key)
		cfg.BaseURL = srv.URL + "/v1"
		return gpt.NewClientWithConfig(cfg)
	}
	defer func() { newClient = old }()

	t.Setenv("OPENAI_API_KEY", "k")
	ch := Stream("hi")
	var out string
	for tok := range ch {
		out += tok
	}
	if out != "hi there" {
		t.Fatalf("got %q", out)
	}
}

func TestStreamNoKey(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	ch := Stream("hi")
	if _, ok := <-ch; ok {
		t.Fatalf("expected closed channel")
	}
}
