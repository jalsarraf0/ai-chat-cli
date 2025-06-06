#!/usr/bin/env bash
# Copyright 2025 The ai-chat-cli Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

command -v go >/dev/null || { echo "❌ 'go' not on PATH"; exit 1; }
[[ $(go version) =~ go1\.24 ]] || { echo "❌ Go 1.24.x required"; exit 1; }

export PATH="$PWD/offline-bin:$PATH"
mkdir -p offline-bin

curl -sSfI https://proxy.golang.org >/dev/null || {
  echo "❌ no network, can't install tools" >&2; exit 1; }

echo "🔧 installing tools"
for pkg in \
  mvdan.cc/gofumpt@latest \
  honnef.co/go/tools/cmd/staticcheck@latest \
  github.com/securego/gosec/v2/cmd/gosec@latest \
  golang.org/x/vuln/cmd/govulncheck@latest \
  github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1; do
    GOFLAGS='-trimpath' go install "$pkg"
done
