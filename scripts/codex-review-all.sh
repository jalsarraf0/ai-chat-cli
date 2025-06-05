#!/usr/bin/env bash
set -euo pipefail

mkdir -p codex/reports
for pkg in $(go list ./...); do
  name=$(echo "$pkg" | tr '/.' '-')
  out="codex/reports/${name}.md"
  echo "reviewing $pkg -> $out" >&2
  openai tools codex review "$pkg" > "$out"
  sleep 1
done
