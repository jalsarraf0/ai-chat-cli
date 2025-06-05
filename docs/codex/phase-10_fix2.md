<!--
Phase 10 - Fail-Fix Patch – slice runner & Windows make
Save as `docs/codex/phase-10_fix2.md`
Generated 2025-06-05
-->

# Phase 10 – **Green Matrix Hot‑fix** 🟢

This patch fixes the two blockers seen:

| Job | Failure | Root cause | Fix |
|-----|---------|------------|-----|
| `unit‑coverage (linux-slice‑X)` | `gotestsum unknown flag: --subset` | old binary in PATH | pin gotestsum ≥ `1.11.0` and fall back to native slice calc |
| `unit‑coverage (windows)` | Chocolatey `make` install fails (non‑admin lock) | our self‑hosted runner lacks `make`; install needs elevation | run tests **without make** on Windows |

---

## 1 • Tool pinning

### `scripts/bootstrap-tools.sh` / `.ps1`

Append **explicit module path** which contains the slice feature:

```bash
gotest.tools/gotestsum/v2/cmd/gotestsum@v1.11.0
```

PowerShell array gets the same entry.

This replaces any distro copy.

---

## 2 • Slice without `--subset` fallback

Replace the unit test command in `ci.yml`:

```bash
IDX="${CASE%%/*}" TOT="${CASE##*/}"
if [[ "$TOT" != "$CASE" ]]; then
  # calculate nth subset using pure bash ➜ go list shards
  PKGS=$(go list ./... | awk -v idx="$IDX" -v tot="$TOT" 'NR % tot == idx')
else
  PKGS="./..."
fi

go test -race -covermode=atomic -coverprofile=coverage.out $PKGS
```

No dependency on gotestsum flags; still parallel‑safe.

---

## 3 • Windows job without **make**

In matrix include key `os: windows` but set `use_make: false`.

```yaml
- name: Lint (Windows direct)
  if: runner.os == 'Windows'
  run: golangci-lint run ./...

- name: Unit (Windows direct)
  if: runner.os == 'Windows'
  run: go test ./... -race -covermode=atomic -coverprofile=coverage.out
```

Remove Chocolatey step entirely.

Coverage artefact upload remains identical.

---

## 4 • Quality job unchanged

The merged coverage already passes once Linux slices complete
(`macos` now excluded from gate as it duplicates Linux).

---

## 5 • Acceptance check

- [ ] `gotestsum -v` prints **v1.11.0** on Linux slices.
- [ ] Linux slices succeed, upload `coverage.out`.
- [ ] Windows job green (no Chocolatey).
- [ ] `quality` merges profiles → **≥ 92 %**.
- [ ] Docs, snapshot & release remain green.

---

MIT © 2025 Jamal Al‑Sarraf
