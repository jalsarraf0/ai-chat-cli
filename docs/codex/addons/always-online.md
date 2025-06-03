This document **augments** the baseline “Codex Rules”.
If any statement here appears to conflict with the base rules, **apply both** – follow the stricter instruction or raise a question.

──────────────────────────────────────────────────────────────────────────────
▍0. Internet Access Policy
──────────────────────────────────────────────────────────────────────────────
1. **Assume continuous Internet connectivity** in all local scripts, helper tools, and CI jobs.
2. Whenever a dependency is missing or out-of-date you **MUST**:

   ```bash
   go get -u ./...
   go mod tidy -e -v
   ```

3. If a third-party CLI (e.g. `gosec`, `gomarkdoc`, `gocovmerge`) is absent, install it via:

   ```bash
   go install <module>@latest
   ```

   Fallback: download its release asset from GitHub Releases.

──────────────────────────────────────────────────────────────────────────────
▍1. Coverage Gate ≥ 83 %
──────────────────────────────────────────────────────────────────────────────
*The pipeline enforces **≥ 80 %**; you must deliver **≥ 83 %** to provide a 3-point safety margin.*

```bash
# If slice testing is NOT active:
go test -race -covermode=atomic -coverprofile=coverage.out ./...

# If slice testing IS active (Phase 1+):
gocovmerge cover*.out > coverage.out
```

```bash
pct=$(go tool cover -func=coverage.out | awk '/^total:/ {gsub("%","");print $3}')
[[ ${pct%.*} -ge 83 ]] || { echo "::error::coverage < 83 %"; exit 1; }
```

──────────────────────────────────────────────────────────────────────────────
▍2. Conflict-Free Pull Requests
──────────────────────────────────────────────────────────────────────────────
*No merge conflicts are tolerated.*

1. Rebase your feature branch onto **`dev`** before opening / updating a PR:

   ```bash
   git fetch origin
   git rebase origin/dev
   ```

2. Resolve any `<<<<`/`>>>>` markers locally, re-run the full test suite, regenerate docs, and only then:

   ```bash
   git push --force-with-lease
   ```

──────────────────────────────────────────────────────────────────────────────
▍3. CI / CD Compliance
──────────────────────────────────────────────────────────────────────────────
*You **MAY edit** the **single** workflow file `.github/workflows/ci.yml` to extend existing jobs (e.g. slice fan-out), **but you MUST NOT add additional workflow files.***

* Stages must still execute in the order:
  **coverage-gate → security-scan → docs**.
* Local helpers belong in `Makefile` or `scripts/` and must respect `GOFLAGS=-mod=vendor` when set.
* All unit tests compile with `-race` and are table-driven (`t.Parallel()`).

──────────────────────────────────────────────────────────────────────────────
▍4. Quick Pre-PR Checklist
──────────────────────────────────────────────────────────────────────────────
- [ ] `go vet ./...` passes cleanly.
- [ ] `go test -race ./...` ≥ 83 % coverage.
- [ ] `go mod tidy -e` leaves no changes.
- [ ] Docs regenerated (`make docs`) with no diff noise.
- [ ] `shellcheck scripts/*.sh` passes (ignore **SC2086** where documented).
- [ ] Branch rebased onto `dev` – **NO CONFLICTS**.
- [ ] Commit message ≤ 72 chars, imperative; PR labelled **codex-generated**.

> **Remember:** Internet is always available — fetch what you need, heal what you break, keep the gate green, and never ship conflicts.
