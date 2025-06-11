#!/usr/bin/env bash
set -euo pipefail
npm install
npm audit
make docs
