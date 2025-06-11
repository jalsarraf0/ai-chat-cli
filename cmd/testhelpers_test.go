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

package cmd

import (
	"context"
	"errors"
	"io"

	"github.com/jalsarraf0/ai-chat-cli/pkg/llm"
)

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, io.ErrClosedPipe }

type stubLLM struct{ models []string }

func (s stubLLM) Completion(context.Context, llm.Request) (llm.Stream, error) { return nil, nil }
func (s stubLLM) ListModels(context.Context) ([]string, error)                { return s.models, nil }

type errLLM struct{}

func (errLLM) Completion(context.Context, llm.Request) (llm.Stream, error) { return nil, nil }
func (errLLM) ListModels(context.Context) ([]string, error)                { return nil, errors.New("fail") }
