<!--
AI-Chat-CLI • Codex Prompt
Phase 9.4 Patch – Upgrade golangci-lint to v2 & unblock CI
Save as docs/codex/phase-9_4-lint-v2.md
Author: Jamal Al-Sarraf <jalsarraf0@gmail.com>
Generated: 2025-06-05
-->

# Phase 9.4 Patch – **golangci-lint v2 Migration** 🛠️
*Goal: restore green matrix, keep quality sweep intact.*

---

## 0 • Root cause
Binary we compile (`github.com/golangci/golangci-lint/cmd/...`) is **v1**; our `.golangci.yml` targets **v2** → config mismatch.

---

## 1 • Fix overview

1. **Install module-path `/v2`** when bootstrapping.
2. **Pin minimum version ≥ `v2.0.1`** (first stable tag).
3. **Update `.golangci.yml`**:
   * Replace **`run.skip-dirs` → `issues.exclude-dirs`**.
   * Remove other deprecated keys.
4. **Verify lint succeeds on all OS legs**.

---

## 2 • Script changes

### 2.1 `scripts/bootstrap-tools.sh`

```bash
# replace old install loop
for pkg in \
  mvdan.cc/gofumpt@latest \
  honnef.co/go/tools/cmd/staticcheck@latest \
  github.com/securego/gosec/v2/cmd/gosec@latest \
  github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1; do
    GOFLAGS='-trimpath' go install "$pkg"
done
```

### 2.2 `bootstrap-tools.ps1`

Add `'github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1'` to `$Pkgs` array.

---

## 3 • CI YAML patch

Remove the dedicated “Install golangci-lint” step; it’s now handled by bootstrap.

```yaml
# delete:
# - name: Install golangci-lint
#   run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@master
```

Bootstrap will pull v2 binary on every job.

---

## 4 • `.golangci.yml` minimal diff

```yaml
-run:
-  skip-dirs:
-    - examples
+issues:
+  exclude-dirs:
+    - examples
```

_No other config keys change._

---

## 5 • Acceptance checklist

- [ ] `golangci-lint --version` prints `version 2.x`.
- [ ] `golangci-lint run` passes on Linux, macOS, Windows.
- [ ] CI matrix green; coverage & security gates unchanged.
- [ ] `.golangci.yml` free of deprecated options.
- [ ] All commits signed **Jamal Al-Sarraf** and ≤ 72 chars.

---

MIT © 2025 Jamal Al-Sarraf
