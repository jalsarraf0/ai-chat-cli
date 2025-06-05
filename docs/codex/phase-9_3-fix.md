<!--
AI-Chat-CLI • Codex Prompt
Phase 9.3 – Matrix CI Fixes & Security Scan Revival (Go 1.24-pinned)
Save as docs/codex/phase-9_3-fix.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
-->

# Phase 9.3 Prompt – **Make the Matrix Green** ✅
*Goal: all OS legs pass — lint · unit · coverage · security*

---

## 0 • Root causes observed

| OS | Error | Fix |
|----|-------|-----|
| **macOS** | `golangci-lint: No such file or directory` | Add dedicated step to compile **golangci‑lint HEAD** with Go 1.24. |
| **Windows** | Bash-style `[[ … ]]` in PowerShell → parser error | Split **GOMAXPROCS** export into OS‑specific steps using the correct shell. |
| **Security scan** | Stage missing after matrix refactor | Re‑add `make security-scan` step (runs `gosec`) in every matrix leg. |

---

## 1 • CI YAML patch (`.github/workflows/ci.yml`)

```yaml
jobs:
  test:
    name: unit-coverage (${{ matrix.os }})
    strategy:
      fail-fast: false
      matrix:
        include:
          - os: linux
            runs-on: [self-hosted, linux]
            shell: bash
          - os: macos
            runs-on: macos-latest
            shell: bash
          - os: windows
            runs-on: windows-latest
            shell: pwsh

    runs-on: ${{ matrix.runs-on }}
    defaults:
      run:
        shell: ${{ matrix.shell }}

    steps:
      - uses: actions/checkout@v4

      # 1. Setup Go 1.24.x so 'go' is on PATH 🚀
      - uses: actions/setup-go@v5
        with:
          go-version: 1.24.x
          cache: true

      # 2. Install golangci-lint @ HEAD (compile w/ 1.24)
      - name: Install golangci-lint
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@master
          golangci-lint --version

      # 3. Bootstrap other critical tools (gofumpt · staticcheck · gosec)
      - name: Bootstrap critical tools
        if: runner.os != 'Windows'
        run: ./scripts/bootstrap-tools.sh install_tools
      - name: Bootstrap critical tools (Windows)
        if: runner.os == 'Windows'
        shell: pwsh
        run: ./scripts/bootstrap-tools.ps1

      # 4. Set GOMAXPROCS per‑OS
      - name: Export GOMAXPROCS (Linux/macOS)
        if: runner.os != 'Windows'
        run: |
          echo "GOMAXPROCS=$(getconf _NPROCESSORS_ONLN 2>/dev/null || sysctl -n hw.ncpu)" >> $GITHUB_ENV
      - name: Export GOMAXPROCS (Windows)
        if: runner.os == 'Windows'
        shell: pwsh
        run: |
          echo "GOMAXPROCS=$Env:NUMBER_OF_PROCESSORS" >> $Env:GITHUB_ENV

      # 5. Lint 🧹
      - name: Lint
        run: make lint

      # 6. Unit tests + coverage gate 🧪
      - name: Unit tests
        run: make unit
      - name: Coverage gate
        run: make coverage-gate

      # 7. Security scan 🔒
      - name: Security scan (gosec)
        run: make security-scan
```

### Down‑stream jobs

```yaml
  docs:
    needs: test          # waits for **all** matrix legs
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: make docs

  # snapshot / release jobs remain unchanged
```

---

## 2 • Makefile – ensure security target exists

```make
.PHONY: security-scan
security-scan:
gosec ./...
```

(Add only if not already present.)

---

## 3 • README badge

Keep a single badge:

```
[![CI (Linux·macOS·Windows)](https://github.com/<ORG>/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](...)
```

---

## 4 • Acceptance checklist

- [ ] **golangci‑lint** compiles & runs on macOS and Windows (no “not found”).
- [ ] Matrix legs export **GOMAXPROCS** without shell errors.
- [ ] `make unit` passes **≥ 92 %** coverage in all legs.
- [ ] `make security-scan` (gosec) succeeds with **0 critical findings**.
- [ ] Workflow green; embed‑drift / docs / snapshot follow.
- [ ] README badge reflects the matrix status.
- [ ] No extra workflow files; diff limited to CI YAML, README, Makefile.
- [ ] Commits signed **Jamal Al‑Sarraf** and ≤ 72 chars.

---

MIT © 2025 Jamal Al‑Sarraf
