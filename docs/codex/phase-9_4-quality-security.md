<!--
AI-Chat-CLI • Codex Prompt
Phase 9.4 – Quality Sweep & Enhanced Security Scan (Go 1.24-pinned)
Save this file as docs/codex/phase-9_4-quality-security.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
-->

# Phase 9.4 Prompt – **Repo‑wide Quality & Security Pass** 🦩🔒
_Target: keep velocity high while eliminating hidden risks._

---

## 0 • Immutable rule
> The entire project **must compile & test with Go 1.24.x** on Linux, macOS, and Windows.

---

## 1 • Objectives

| Pillar | Deliverable |
|--------|-------------|
| **Lint & Style** | Zero issues from `golangci-lint run` with the **full** ruleset (no `--fast`). |
| **Static Analysis** | `staticcheck ./...` integrated into CI; no SA errors. |
| **Security** | 1️⃣ `gosec ./...` (high‑signal rules)<br>2️⃣ `govulncheck ./...` (Go 1.24 builtin)<br>Both stages fail the build on any **medium+ severity** finding. |
| **Docs** | All exported symbols have godoc comments; `mkdocs build` clean. |
| **Dependencies** | `go mod tidy --go=1.24` + `go mod verify` run in CI. |

---

## 2 • Tooling bootstrap additions

Append to **`scripts/bootstrap-tools.sh`** lists:

```bash
  mvdan.cc/gofumpt@latest
  honnef.co/go/tools/cmd/staticcheck@latest
  github.com/securego/gosec/v2/cmd/gosec@latest
  golang.org/x/vuln/cmd/govulncheck@latest      # NEW
```

PowerShell script updated analogously.

---

## 3 • `.golangci.yml` upgrades

* Enable high‑signal linters: `revive`, `gocritic`, `misspell`, `nestif`.
* Set `deadline: 5m`.
* Exclude‑path for `examples/**` so sample plugins don’t block.

---

## 4 • CI workflow changes (`ci.yml`)

Add an **additional job** `quality` that runs after the matrix:

```yaml
  quality:
    needs: test                   # waits for all OS legs
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: 1.24.x
          cache: true

      - name: Bootstrap tools
        run: ./scripts/bootstrap-tools.sh install_tools

      - name: Lint (all linters)
        run: golangci-lint run ./...

      - name: Staticcheck
        run: staticcheck ./...

      - name: Security scan – gosec
        run: gosec ./...

      - name: Security scan – govulncheck
        run: govulncheck ./...
```

`docs`, `snapshot`, `release` now depend on `quality` instead of `test`.

---

## 5 • Makefile helpers

```make
.PHONY: lint-all static security
lint-all: ; golangci-lint run ./...
static:   ; staticcheck ./...
security: ; gosec ./... && govulncheck ./...
```

---

## 6 • Documentation

* **docs/security.md** – outline dual scan (`gosec` + `govulncheck`), how to suppress with comments.
* **CHANGELOG** – add `## [0.9.4] – YYYY‑MM‑DD` entry.
* **README** – badge “Security Scan 🗡️ 100 % clean”.

---

## 7 • Acceptance checklist

- [ ] `golangci-lint run` reports **0 issues** across repo.
- [ ] `staticcheck` passes.
- [ ] `gosec` & `govulncheck` report **0 medium+ findings**.
- [ ] Coverage gate still **≥ 92 %**.
- [ ] CI workflow green: `test → quality → docs`.
- [ ] `go mod tidy` produces no diff.
- [ ] Docs site builds without warnings.
- [ ] Commits signed **Jamal Al‑Sarraf** and ≤ 72 chars.

---

MIT © 2025 Jamal Al‑Sarraf
