# Codex Prompt - Always-Online & SUPREME-Release Add-On
*(Save or overwrite as docs/codex/addons/always-online.md and include it in every change-set or PR.)*

This document **augments** the baseline Codex Rules.
If any statement here conflicts with the base rules, obey the *stricter* instruction or raise a question.

---
## 0 · Internet Access Policy \ud83c\udf10
Our self-hosted runners ship with **pre-baked, version-locked binaries**
\u2026patched text\u2026
6. **New CLI tool** needed?  Open an issue. After approval, bake it into
   `scripts/preflight-*.sh`; never inline-install inside workflow YAML.

---
## \▍1. Coverage Gate \u2265\u200a93\u200a% \ud83d\udcc8
Pipeline minimum is 90\u200a%; you **must deliver \u2265\u200a93\u200a%** for a 3\u2011pp buffer.

```bash
go test -race -covermode=atomic -coverprofile=coverage.out ./...
pct=$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $3}')
[[ ${pct%.*} -ge 93 ]] || { echo "::error::coverage < 93\u200a%"; exit 1; }
```

For slice testing, merge first:

```bash
gocovmerge cover*.out > coverage.out
```

---
## \▍2. Full Security Suite \ud83d\udee1\ufe0f
Every PR **must** pass **all** checks *with zero HIGH / CRITICAL findings*.

| Tool | Command |
|------|---------|
| **Gosec** | `gosec ./...` |
| **Govulncheck** | `govulncheck ./...` |
| **Trivy (FS)** | `trivy fs --exit-code 1 --severity HIGH,CRITICAL .` |
| **License Audit** | `addlicense -check $(go list -f '{{.Dir}}' ./...)` |

## §3 · Immutable CI/CD Pipelines \ud83d\udd12
No file under .github/workflows/**, .azure-pipelines/**, or
.circleci/** may be added, removed, or edited in a pull-request
without an issue labelled ci-approval and an approving review from
the Maintain CI GitHub team.

Any PR that touches those paths without approval must be closed or
reverted immediately.

If a pipeline change is essential (e.g., security patch), open an
issue, tag the CI maintainers, and wait for sign-off before pushing
code.

---
## \▍4. CI / CD Flexibility \u2699\ufe0f
*You **may** create or modify workflow files*, provided:

1. **Gate order** stays **unit \u2192 quality \u2192 security \u2192 build \u2192 snapshot \u2192 release \u2192 docs**.
2. Unit tests use `-race`, are table-driven (`t.Parallel()`).
3. Jobs honour `GOFLAGS="-mod=vendor"` when set.
4. Self-hosted **Linux & Windows** gates finish **before** `macos-latest`.
5. Release jobs trigger only on `v*` tags and sign artefacts with `cosign`.

---
## \▍5. Conflict-Free Pull Requests \ud83d\udd00
1. Rebase your branch on **`dev`** before each push:

   ```bash
   git fetch origin
   git rebase origin/dev
   ```

2. Resolve conflicts, rerun full suite, regenerate docs, then:

   ```bash
   git push --force-with-lease
   ```

---
## \▍6. Quick Pre-PR Checklist \u2705
- [ ] `go vet ./...` passes.
- [ ] `go test -race ./...` shows **\u2265\u200a93\u200a%** coverage.
- [ ] Security suite (Gosec\u200a+\u200aGovulncheck\u200a+\u200aTrivy) green.
- [ ] `go mod tidy -e` leaves no diff.
- [ ] Docs regenerated (`make docs`) with no noise.
- [ ] `shellcheck scripts/*.sh` passes (ignore **SC2086** where noted).
- [ ] Branch rebased \u2014 **no conflicts**.
- [ ] Commit message \u2264\u200a72 chars, imperative; PR labelled **codex-generated**.

> **Always stay online, hit \u2265\u200a93\u200a% coverage, pass every security gate, and ship cleanly.**
