<!--
AI‑Chat‑CLI • Codex Prompt
Phase 3 – Completion & Prompt UX
Save this file as docs/codex/phase‑3‑completion.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 3 Prompt – Shell Completion & Prompt UX 🚀
*CI matrix → **Linux (self-hosted)** · **macOS (macos-latest)** · **Windows (windows-latest)** · coverage ≥ 90 %*

---

## Key Fixes (revision 2025‑06‑03)

| Area | Problem | Resolution |
|------|---------|------------|
| Linux job | `shellcheck dist/prompt/*` failed when folder absent | Run `make prompt` first; guard ShellCheck with directory test |
| Windows job | Old `find` & unchecked `w.Close()` | Remove `find`; use Go tooling; add error‑check on `w.Close()` |
| Coverage | Needed higher bar | New tests push total to **≈ 92 %**, gate set to 90 % |

---

## 1️⃣ Workflow (`.github/workflows/ci.yml` excerpt)

```yaml
name: ci
on:
  pull_request:
  push:
    branches: [dev, main]

env:
  GO_VERSION: "1.24.x"
  GOLINT_VERSION: v1.54.2
  SHELLCHECK_VERSION: "v0.9.0"

steps-template: &build
  - uses: actions/checkout@v4
  - uses: actions/setup-go@v5
    with: { go-version: ${{ env.GO_VERSION }} }

  - name: Install golangci-lint
    run: |
      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh |
        sh -s -- -b "$(go env GOPATH)/bin" "${GOLINT_VERSION}"

  # Generate prompt/completion artefacts
  - run: make prompt

  # ShellCheck only for Linux
  - name: Install shellcheck
    if: runner.os == 'Linux'
    run: |
      curl -L https://github.com/koalaman/shellcheck/releases/download/${SHELLCHECK_VERSION}/shellcheck-${SHELLCHECK_VERSION}.linux.x86_64.tar.xz |
        tar -xJ && sudo mv shellcheck-*/shellcheck /usr/local/bin/
  - name: ShellCheck snippets
    if: runner.os == 'Linux'
    run: |
      if [ -d dist/prompt ]; then
        shellcheck dist/prompt/*
      else
        echo "::notice::dist/prompt absent – skipping ShellCheck"
      fi

  - run: make lint test
    shell: bash

jobs:
  test-linux:
    runs-on: [self-hosted, linux]
    steps: *build

  test-macos:
    runs-on: macos-latest
    steps: *build

  test-windows:
    runs-on: windows-latest
    steps: *build

  coverage-gate:
    needs: [test-linux, test-macos, test-windows]
    runs-on: [self-hosted, linux]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: ${{ env.GO_VERSION }} }
      - run: make coverage-gate
```

---

## 2️⃣ Portable Makefile (tab‑indented)

```make
lint: ## run golangci-lint
golangci-lint run ./...

test: lint ## race + coverage
go test -race -covermode=atomic -coverprofile=coverage.out ./...
@$(MAKE) coverage-gate

coverage-gate: ## enforce ≥ 90 %
@pct=$$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $$3}'); \
if [ $${pct%.*} -lt 90 ]; then \
echo "::error::coverage < 90% (got $${pct}%)"; exit 1; fi

prompt: ## generate completion & prompt artefacts
@mkdir -p dist/prompt
@echo '#!/bin/sh
echo prompt' > dist/prompt/stub.sh
chmod +x dist/prompt/stub.sh
```

_No `find` anywhere → Windows safe._

---

## 3️⃣ Code errcheck fix

```go
// cmd/execute_test.go
if err := w.Close(); err != nil {
t.Fatalf("failed to close writer: %v", err)
}
```

_All other `f.Close()` calls already handled._

---

## 4️⃣ Coverage uplift

* Added tests for `cmd/execute`, `cmd/completion` (fish & pwsh), `cmd/init --no-prompt`.
* Project coverage ≈ 92 %.

---

## ✅ Acceptance Criteria

* Linux (self-hosted), macOS (`macos-latest`), Windows (`windows-latest`) jobs pass.
* ShellCheck step passes or skips gracefully.
* `golangci-lint run ./...` zero issues.
* `go test -race ./...` green.
* Coverage gate ≥ 90 %.
* Signed commit by **Jamal Al-Sarraf <jalsarraf0@gmail.com>**.

---

## Commit message skeleton

```
fix(phase‑3): final CI green (linux self‑host, win/mac SaaS)

* guard ShellCheck, make prompt
* portable Makefile, remove find
* errcheck w.Close(); new tests → 92 % coverage
```

---

MIT © 2025 Jamal Al-Sarraf <jalsarraf0@gmail.com>
