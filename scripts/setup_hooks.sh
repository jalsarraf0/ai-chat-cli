#!/usr/bin/env bash
set -Eeuo pipefail
root=$(git rev-parse --show-toplevel)
ln -sf "$root/scripts/pre_commit.sh" "$root/.git/hooks/pre-commit"
