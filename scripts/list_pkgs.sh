#!/usr/bin/env bash
set -Eeuo pipefail
go list ./... | grep -v /vendor/
