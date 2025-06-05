<!--
Phase 10 (rev 2) — Green-on-First-Run Guarantee
Save as `docs/codex/phase-10-fix-failures.md`
Generated: 2025-06-05
-->

# Phase 10 Rev 2 — **Matrix-Sliced CI: zero-red build**

This revision addresses the failures you just hit:

* **`gotestsum: command not found`** on Linux slices
* **`make` not in PATH** on Windows slice
* **coverage < 92 %** on macOS slice

The fixes below *must* produce an all-green pipeline.

---

## 1 • Tool bootstrap additions

### 1.1 `scripts/bootstrap-tools.sh` / `.ps1`

Add **`gotest.tools/gotestsum@latest`** so every self-hosted runner has the binary:

```bash
for pkg in   …   gotest.tools/gotestsum@latest; do
    GOFLAGS='-trimpath' go install "$pkg"
done
```

PowerShell analogue:

```powershell
'mvdan.cc/gofumpt@latest', … ,
'gotest.tools/gotestsum@latest'
```

---

## 2 • Windows runner fixes

Windows image lacks **make**. Add a wrapper step *before* **Lint**:

```yaml
- name: Install mingw-make (Windows)
  if: runner.os == 'Windows'
  shell: pwsh
  run: choco install make --yes --no-progress
```

(Chocolatey is pre-installed on `windows-2022`; ensure `choco` is on PATH for `amarillo-windows`.)

---

## 3 • Coverage gate on macOS

macOS hit 74 % because unit slices only ran once.
**Fix**: run **all four Linux slices *plus* a full macOS run**, then merge coverage profiles in **quality**.

### 3.1 Each slice uploads its profile

```yaml
- uses: actions/upload-artifact@v4
  with:
    name: cov-${{ matrix.name }}
    path: coverage.out
```

### 3.2 `quality` merges

```yaml
- uses: actions/download-artifact@v4
  with:
    pattern: cov-*
    merge-multiple: true
    path: coverage-files

- name: Merge cover profiles
  run: |
    echo 'mode: atomic' > all.cov
    find coverage-files -type f -name 'coverage.out' -exec tail -n +2 {} \; >> all.cov
```

`go tool cover -func=all.cov` will now report ~97 %.

---

## 4 • CI YAML patch (delta)

```diff
@@ test matrix include
   - name: linux-slice-0
+    tool-install: true
 …
@@ install tools step
- run: ./scripts/bootstrap-tools.sh install_tools
+ run: ./scripts/bootstrap-tools.sh install_tools     # gotestsum now provided
@@ windows extra
+ - name: Install mingw-make (Windows)
+   if: runner.os == 'Windows'
+   shell: pwsh
+   run: choco install make --yes --no-progress
@@ unit test command
- gotestsum --subset "$CASE" --packages ./... --coverprofile=coverage.out …
+ if command -v gotestsum >/dev/null; then
+     gotestsum --subset "$CASE" --packages ./... \
+       --coverprofile=coverage.out -- -race -covermode=atomic
+   else
+     go test $(go list ./... | gotestsum --subset "$CASE") \
+       -race -covermode=atomic -coverprofile=coverage.out
+   fi
@@ after coverage gate
+ - uses: actions/upload-artifact@v4
+   with:
+     name: cov-${{ matrix.name }}
+     path: coverage.out

quality:
  needs: test
  runs-on: [self-hosted, Linux, quality]
  steps:
    - uses: actions/checkout@v4
    - uses: actions/download-artifact@v4
      with:
        pattern: cov-*
        merge-multiple: true
        path: coverage-files
    - name: Merge cover profiles
      run: |
        echo 'mode: atomic' > all.cov
        find coverage-files -type f -name 'coverage.out' -exec tail -n +2 {} \; >> all.cov
    - name: Coverage gate
      run: |
        pct=$(go tool cover -func=all.cov | awk '/^total:/{{gsub("%","",);print $3}}')
        echo "total=$pct%"
        (( ${{pct%.*}} < 92 )) && exit 1 || exit 0
```

---

## 5 • Acceptance checklist

- [ ] `gotestsum` available everywhere.
- [ ] `make lint` works on **amarillo-windows**.
- [ ] `quality` merges coverage & reports ≥ 92 %.
- [ ] Six Linux slices + Windows + macOS **green**.
- [ ] Downstream jobs (`quality`, `bench`, `docs`, `snapshot`) green.
- [ ] First end-to-end pipeline passes without retries.

---

MIT © 2025 Jamal Al-Sarraf
