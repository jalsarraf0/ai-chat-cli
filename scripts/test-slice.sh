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

# Run unit tests slice with coverage
# shellcheck disable=SC2086
set -euo pipefail

GOFLAGS=${GOFLAGS:-}

case ${CASE:-1/1} in
  */*)
    IFS=/ read -r index total <<<"$CASE"
    ;;
  *)
    index=1
    total=1
    ;;
esac

mapfile -t packages < <(go list ./...)
count=${#packages[@]}
size=$(( (count + total - 1) / total ))
start=$(( (index - 1) * size ))
selected=("${packages[@]:start:size}")

go test $GOFLAGS -race -covermode=atomic -coverprofile=coverage.out -tags unit "${selected[@]}"
