<!--
AI\u2011Chat\u2011CLI \u2022 Codex Prompt
Phase\u00a08 \u2013 Cross\u2011Compilation & Release Automation
Save this file as docs/codex/phase\u20118\u2011release.md
Author: Jamal\u00a0Al\u2011Sarraf <jalsarraf0@gmail.com>
-->

# Phase\u00a08 Prompt \u2013 Cross\u2011Compilation & Automated Releases \ud83d\ude80
*Runner matrix: **Linux\u202f[self\u2011hosted,\u2011linux] \xb7 Windows\u202f[self\u2011hosted,\u2011windows] \xb7 macOS\u202f(macos\u2011latest)** \u2014 coverage gate **\u2265\u202f93\u202f%***

---

## Deliverables

1. **`goreleaser.yml`** \u2013 multi\u2011arch builds (linux/amd64,arm64; windows/amd64; darwin/amd64,arm64), `CGO_ENABLED=0`, archives include completions & themes, SHA256 + SBOM.
2. **Cosign signing** \u2013 use repo secrets `COSIGN_PRIVATE_KEY`, `COSIGN_PASSWORD`.
3. **Docker images** \u2013 `ghcr.io/jalsarraf0/ai-chat-cli:{version,latest}` produced via Goreleaser.
4. **Homebrew formula** \u2013 auto\u2011tap `jalsarraf0/homebrew-ai-chat-cli`.
5. **Workflows**
   * **Snapshot** \u2013 extend CI: `make snapshot` (`goreleaser build --snapshot`).
   * **Release** \u2013 new `release.yml`, trigger on `v*.*.*` tag, runs `goreleaser release --clean --sign` on Linux self\u2011hosted.
6. **Makefile targets**
   ```make
   snapshot:
    \tgoreleaser build --snapshot --clean
   release:
    \tgoreleaser release --clean --sign
   ```
7. **Version stamping** \u2013 ldflags inject version/commit/date; `ai-chat version` prints them.
8. **Tests to keep \u2265\u202f93\u202f% coverage**
   * `version_test.go` \u2013 ensure stamped info non\u2011empty.
   * `sbom_test.go` \u2013 extract snapshot archive, assert `sbom.json` present.
9. **macOS cache fix** \u2013 clean `${{ env.CACHE_DIR }}` before cache restore.
10. **Docs** \u2013 `docs/releases.md`, README badge for latest release & Docker pull.

---

## Key CI snippets

### Set full cores on Windows

```yaml
- name: Set GOMAXPROCS
  if: runner.os == 'Windows'
  run: echo "GOMAXPROCS=$Env:NUMBER_OF_PROCESSORS" >> $Env:GITHUB_ENV
```

### Release workflow (linux self\u2011hosted)

```yaml
on: { push: { tags: ["v*.*.*"] } }
jobs:
  release:
    runs-on: [self-hosted, linux]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: 1.24.x }
      - run: goreleaser release --clean --sign
        env:
          COSIGN_PRIVATE_KEY: ${{ secrets.COSIGN_PRIVATE_KEY }}
          COSIGN_PASSWORD:    ${{ secrets.COSIGN_PASSWORD }}
```

---

## Acceptance checklist

- Snapshot artifacts attach to PRs (Linux, Windows, macOS bins + SBOM & checksums).
- Tag `v0.1.0` produces GitHub Release, signed archives, Homebrew formula, Docker images.
- `cosign verify` passes.
- Coverage \u2265\u202f93\u202f%.
- Docs merged; commit signed **Jamal\u00a0Al\u2011Sarraf**.

---

MIT \u00a9\u00a02025\u00a0Jamal\u00a0Al\u2011Sarraf
