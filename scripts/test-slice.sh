#!/usr/bin/env bash
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
