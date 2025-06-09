#!/usr/bin/env bash
set -euo pipefail

find . -name '*_test.go' -print0 | while IFS= read -r -d '' file; do
  dir=$(dirname "$file")
  target=$([[ "$dir" == "." ]] && echo "main_test" || echo "cmd_test")
  pkg_line=$(grep -n '^package ' "$file" | head -n1 | cut -d: -f1)
  current=$(awk "NR==$pkg_line {print \$2}" "$file")
  [[ "$current" == "$target" ]] && continue
  echo "Fixing $file → $target"
  sed -i "${pkg_line}s/^package .*/package ${target}/" "$file"
done

"$HOME/go/bin/golangci-lint" run --timeout 5m
