#!/usr/bin/env bash
# Usage: SLICES=7 SLICE_INDEX=3 scripts/test_slice.sh
set -Eeuo pipefail
total=${SLICES:-7}
idx=${SLICE_INDEX:-1}
mapfile -t pkgs < <(scripts/list_pkgs.sh)
sel=()
for i in "${!pkgs[@]}"; do
  (( i % total + 1 == idx )) && sel+=("${pkgs[i]}")
done
# shellcheck disable=SC2086
go test -race -covermode=atomic -coverprofile "cover${idx}.out" "${sel[@]}"
