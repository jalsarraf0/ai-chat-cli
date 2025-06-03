
# AI Chat CLI

## Parallel test slices
CI runs tests in 7 parallel slices (env `SLICES=7`).
Locally you can run `make test-slice` to replicate the sharding.

# AI-Chat-CLI

![CI](https://github.com/jalsarraf0/ai-chat-cli/actions/workflows/ci.yml/badge.svg)

This project provides a minimal CLI built with Go and Cobra.

## Getting started

Prerequisites: Go >= 1.24.3

Clone the repo and build:

```bash
go build ./cmd/ai-chat
```

## Build status

The CI pipeline ensures tests run with race detection and coverage above 83%.

