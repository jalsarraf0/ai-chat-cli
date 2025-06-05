<!--
AI-Chat-CLI • Codex Prompt
Phase 10 – Matrix Slicing ✂️ + Full-Repo Code-Review Sweep
Save as: docs/codex/phase-10-matrix-slice-and-review.md
Generated: 2025-06-05
-->

# Phase 10 — **Matrix-Sliced CI & Automated Code Review** 🚀🧐

> Deliver a dramatically faster pipeline while raising code quality.
> **All gates must be green on the *first* CI run.**

---

## 0 • Immutable constraints

1. **Go 1.24.x** across Linux, Windows, macOS.
2. Coverage **≥ 92 %** on every slice.
3. `quality` stage runs **only** on the dedicated runner `amarillo-runner-01 (self-hosted, Linux, quality)`.
4. No macOS self-hosted runner yet — keep GitHub-hosted for that leg.

---

## 1 • CI workflow design (`.github/workflows/ci.yml`)

### 1.1 Test stage – matrix slices

| Slice | Runs-on | Shell | Purpose |
|-------|---------|-------|---------|
| `0/4` | `[self-hosted, Linux, docker]` | bash | first quarter of pkgs |
| `1/4` | same | bash | second quarter |
| `2/4` | same | bash | third quarter |
| `3/4` | same | bash | fourth quarter |
| `win` | `[self-hosted, Windows, quality]` | pwsh | full suite on Windows |
| `mac` | `macos-latest` | bash | full suite on macOS |

```diff
strategy:
  matrix:
    include:
      - name: linux-slice-0
        slice: 0/4
        runs-on: [self-hosted, Linux, docker]
        shell: bash
      …
      - name: windows
        slice: full
        runs-on: [self-hosted, Windows, quality]
        shell: pwsh
      - name: macos
        slice: full
        runs-on: macos-latest
        shell: bash
```

In the test step:

```bash
CASE="${ matrix.slice }"
if [[ "$CASE" == */* ]]; then
  gotestsum --subset "$CASE" --packages ./... --coverprofile=coverage.out -- -race -covermode=atomic
else
  go test -race -covermode=atomic -coverprofile=coverage.out ./...
fi
```

### 1.2 Quality stage (unchanged)

```yaml
quality:
  needs: test
  runs-on: [self-hosted, Linux, quality]
```

### 1.3 Bench stage

Runs only after *quality* on one Linux docker slice.

```yaml
bench:
  needs: quality
  runs-on: [self-hosted, Linux, docker]
  steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v5
      with:
        go-version: 1.24.x
    - run: make bench-json      # outputs bench.json
    - uses: actions/upload-artifact@v4
      with:
        name: bench
        path: bench.json
```

---

## 2 • Makefile additions

```make
.PHONY: bench bench-json

bench: ; go test -run=^$ -bench=. ./...
bench-json:
go test -run=^$ -bench=. -benchmem -count=3 -json ./... > bench.json
```

---

## 3 • Automated code-review sweep

1. Add **`scripts/codex-review-all.sh`**:
   loops over packages, invokes GPT-4o via `openai tools codex review`
   and writes suggestions to `codex/reports/*.md`.
2. Create **GitHub Action `codex-review`** (manual dispatch) that runs the
   script, commits the reports on a new branch, and opens a draft PR.
3. Triage the PRs, group changes ≤ 300 LoC each, label `phase-10-review`.

---

## 4 • Documentation

* `CONTRIBUTING.md` with guidelines:
  `gofumpt`, `golangci-lint run`, signed commits, 72-char max summary.
* Update **README** badge to show *bench* artefact link.

---

## 5 • Acceptance checklist

- [ ] Matrix slices complete in parallel; wall time ≤ 50 % of previous run.
- [ ] **All gates green on first CI attempt** (test, quality, bench, docs).
- [ ] Coverage aggregated ≥ 92 %.
- [ ] `bench.json` artefact present and downloadable in run summary.
- [ ] At least one Codex review PR merged; remaining suggestions ticketed.
- [ ] CHANGELOG entry `## [0.10.0] – YYYY-MM-DD`.
- [ ] Tag `v0.10.0` + snapshot release created.

---

MIT © 2025 Jamal Al-Sarraf
