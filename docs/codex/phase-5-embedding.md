<!--
AI‑Chat‑CLI • Codex Prompt
Phase 5 – Offline Asset Embedding
Save this file as docs/codex/phase‑5‑embedding.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 5 Prompt – Offline Asset Embedding 📦
*CI matrix: **Linux [self‑hosted, linux] · Windows [self‑hosted, windows] · macOS (macos‑latest)**
Coverage gate: **≥ 90 %** (target ~92 %).*

---

## 🎯 Deliverables (must all be implemented)

| # | Item | Precise specification |
|---|------|-----------------------|
| 1 | **Embedded file‑system** | Create `internal/assets/fs.go`:<br>`//go:embed templates/*.tmpl themes/*.json` bundles:<br>  • `templates/default.tmpl` & `templates/system.tmpl`<br>  • `themes/light.json` & `themes/dark.json` |
| 2 | **Utility package `pkg/embedutil`** | Functions:<br>`List() []string` – returns sorted logical names,<br>`Read(name string) ([]byte, error)`,<br>`MustText(name string) string` – panics on error. |
| 3 | **CLI command `ai-chat assets`** | Cobra sub‑commands:<br>`list` — prints names one per line.<br>`cat <name>` — writes file contents to stdout.<br>`export <name> <file>` — creates parent dirs, writes bytes; rejects overwrite unless `--force`. |
| 4 | **Drift guard** | Script `scripts/embedgen.go` updates `internal/assets/assets_gen.go` with SHA‑256 table.<br>`make embed-check` runs the script **and** fails if `git diff --quiet` is non‑zero. |
| 5 | **CI integration** | New job **`embed-drift`** after `security-scan`:<br>runs on `[self-hosted, linux]` and executes `make embed-check`. |
| 6 | **Documentation** | Add `docs/assets.md` explaining embedded files, override pattern (`--config templatesDir`). Update README with an “Embedded Assets” section + command examples. |
| 7 | **Test coverage ≥ 90 %** | New tests:<br>• `embedutil_test.go` – covers 100 % of utility functions.<br>• `cmd/assets_test.go` – exercise all sub‑commands via Cobra’s test harness.<br>This lifts global coverage to ≈ 92 %. |

---

## 🔧 Reference snippets

### internal/assets/fs.go

```go
package assets

import "embed"

// FS embeds default templates and colour themes.
 //go:embed templates/*.tmpl themes/*.json
var FS embed.FS
```

### pkg/embedutil/embedutil.go

```go
package embedutil

import (
    "bytes"
    "io/fs"
    "sort"

    "github.com/jalsarraf0/ai-chat-cli/internal/assets"
)

func List() []string {
    var names []string
    fs.WalkDir(assets.FS, ".", func(path string, d fs.DirEntry, err error) error {
        if !d.IsDir() {
            names = append(names, path)
        }
        return nil
    })
    sort.Strings(names)
    return names
}

func Read(name string) ([]byte, error) { return assets.FS.ReadFile(name) }

func MustText(name string) string {
    data, err := Read(name)
    if err != nil {
        panic(err)
    }
    return string(bytes.TrimSpace(data))
}
```

### Makefile additions (tabs preserved)

```make
embed-check: ## verify embedded FS is up to date
    go run scripts/embedgen.go
    @git diff --quiet internal/assets || (echo "::error::embed drift"; exit 1)
```

`coverage-gate` remains at 90 % from Phase 4.

---

## 🖥️ CI Workflow injection (`.github/workflows/ci.yml`)

```yaml
  embed-drift:
    needs: security-scan
    runs-on: [self-hosted, linux]
    steps:
      - uses: actions/checkout@v4
      - name: Set GOMAXPROCS
        run: echo "GOMAXPROCS=$(nproc)" >> $GITHUB_ENV
      - run: make embed-check
```

_No other workflow labels change._

---

## 🧪 Testing requirements

1. **embedutil_test.go**
   * Verify `List()` returns exactly four known assets.
   * `MustText("templates/default.tmpl")` contains “{{.User}}”.
2. **cmd/assets_test.go**
   * `assets list` output has 4 lines.
   * `assets cat templates/system.tmpl` matches embedded bytes.
   * `assets export themes/light.json <tmp>` writes an identical file; rerun with `--force` to overwrite.

These tests execute on Linux, Windows, and macOS; use `t.TempDir()` for paths. Combined package coverage pushes global ≥ 90 %.

---

## ✅ Acceptance checklist

- CI matrix passes: Linux and Windows self‑hosted, macOS‑latest.
- `make embed-check` passes both locally and in CI.
- `ai-chat assets` command works exactly as specified.
- `golangci-lint`, `shellcheck`, `gosec`, race detector all green.
- README & `docs/assets.md` updated without diff noise.
- Signed commit **Jamal Al‑Sarraf <jalsarraf0@gmail.com>**.

---

## ✉️ Commit message guideline

```
feat(phase‑5): embed templates & themes, add assets command

* internal/assets via go:embed + drift guard
* pkg/embedutil helpers + 100% tests
* ai-chat assets {list,cat,export --force}
* embed-check job in CI, total coverage 92%
```

---

MIT © 2025 Jamal Al‑Sarraf <jalsarraf0@gmail.com>
