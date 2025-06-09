#!/bin/sh
set -e
GOOS="$1"
ARCH="$2"
SRC="$3"
EXT="$4"
mkdir -p "dist/${GOOS}_${ARCH}"
cp "$SRC" "dist/${GOOS}_${ARCH}/ai-chat-cli${EXT}"
