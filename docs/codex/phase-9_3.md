<!--
AI-Chat-CLI • Codex Prompt
Phase 9.3 – CI Matrix Refactor & Unified Status Badge
Save as docs/codex/phase-9_3.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
-->

# Phase 9.3 – **DRY CI via OS Matrix** 🧩✨
*Collapse three OS‑specific jobs into one matrix while preserving all gates.*

---

## 0 • Immutable facts
* Go version is locked to **1.24.x** in every leg.
* The build steps (lint → unit + coverage → security‑scan → docs → snapshot/release) **must stay in that order**.
* Self‑hosted Linux runner label is `[self-hosted, linux]`.

---

## 1 • Deliverables

| Item | Requirement |
|------|-------------|
| **Matrix job** | Replace `test-linux`, `test-macos`, `test-windows` with a single job `test` using a `matrix.os` strategy. |
| **Bootstrap step** | Keep the _Bootstrap critical tools_ step inside the matrix (with Pwsh branch for Windows). |
| **Cache keys** | Use `${{ matrix.os }}` so each OS keeps its own module cache. |
| **Docs job** | Remains separate and depends on **`test`**. |
| **README badge** | Swap three per‑OS badges for **one unified badge** that points to `ci.yml`. |

---

## 2 • Example workflow patch (`.github/workflows/ci.yml`)
```yaml
jobs:
  test:
    name: unit‑coverage (${{ matrix.os }})
    strategy:
      fail-fast: false
      matrix:
        include:
          # self-hosted Linux
          - os: linux
            runs-on: [self-hosted, linux]
            shell: bash
          # GitHub macOS
          - os: macos
            runs-on: macos-latest
            shell: bash
          # GitHub Windows
          - os: windows
            runs-on: windows-latest
            shell: pwsh

    runs-on: ${{ matrix.runs-on }}
    defaults:
      run:
        shell: ${{ matrix.shell }}

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: ${{ env.GO_VERSION }}

      # OS‑specific GOMAXPROCS
      - name: Set GOMAXPROCS
        run: |
          if [[ "${{ matrix.os }}" == "linux" ]]; then
            echo "GOMAXPROCS=$(nproc)" >> $GITHUB_ENV
          elif [[ "${{ matrix.os }}" == "macos" ]]; then
            echo "GOMAXPROCS=$(sysctl -n hw.ncpu)" >> $GITHUB_ENV
          else
            echo "GOMAXPROCS=$Env:NUMBER_OF_PROCESSORS" >> $Env:GITHUB_ENV
          fi

      # 🔧 Bootstrap linters & gosec
      - name: Bootstrap critical tools
        if: runner.os != 'Windows'
        run: ./scripts/bootstrap-tools.sh install_tools
      - name: Bootstrap critical tools (Windows)
        if: runner.os == 'Windows'
        shell: pwsh
        run: ./scripts/bootstrap.ps1

      # 🧹 Lint + tests + coverage gate
      - run: make lint unit

      # 🔒 Security scan
      - run: make security-scan
```

### Docs & downstream jobs
```yaml
  docs:
    needs: test
    runs-on: [self-hosted, linux]
    steps:
      - uses: actions/checkout@v4
      - run: make docs
```
_Snapshot_ and _release_ jobs stay untouched.

---

## 3 • README badge update
```md
[![CI (Linux · macOS · Windows)](https://github.com/<ORG>/ai-chat-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/<ORG>/ai-chat-cli/actions/workflows/ci.yml)
```
Remove the three obsolete per‑OS badges.

---

## 4 • Makefile (reference)
No changes, but ensure `make unit` already calls coverage‑gate and that
`make security-scan` exists (added in Phase 9.2).

---

## 5 • Acceptance checklist
- [ ] `.github/workflows/ci.yml` contains **one** matrix job `test` covering all three OSes.
- [ ] Self‑hosted Linux runner label unchanged.
- [ ] All matrix legs print **go 1.24.x** and pass: lint ✓, unit + coverage ≥ 92 %, security‑scan ✓.
- [ ] `docs` job runs after the matrix.
- [ ] README shows a single status badge.
- [ ] No other workflow files added; snapshot/release unaffected.
- [ ] Commits signed **Jamal Al‑Sarraf**, ≤ 72 chars.

---

MIT © 2025 Jamal Al‑Sarraf
