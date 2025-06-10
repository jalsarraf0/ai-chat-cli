<!--
AI‑Chat‑CLI • Codex Prompt
Phase 7 – LLM Integration (Coverage Fix)
Save this file as docs/codex/phase‑7‑llm.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 7 Prompt – OpenAI Adapter & Coverage Boost 🤖
*Runner matrix: **Linux [self‑hosted] · Windows [self‑hosted] · macOS (macos‑latest)** — coverage gate **≥ 90 %***

---

## Mandatory tasks

1. **Add high‑coverage unit tests** (`//go:build unit`) for `pkg/llm/openai` covering 5 edge cases via `httptest.Server`.
2. **Refactor openai client**: inject `sleep func(time.Duration)` for retry; production uses `time.Sleep`, tests use no‑op.
3. **Create helper** `internal/testhttp/roundtrip.go` for transport injection.
4. **Add `embedutil` panic test**.
5. **Split Make targets**
   ```make
   unit:
go test -race -covermode=atomic -coverprofile=coverage.out -tags unit ./...
   test: lint unit
   ```
6. **macOS cache fix**: pre‑clean `${{ env.CACHE_DIR }}` before `actions/cache@v3` restore.
7. **Docs**: update `docs/llm.md` with “Unit vs Live” test section.
8. **Gate**: total coverage ≥ 90 % on all OSes; live tests remain opt‑in (`OPENAI_API_KEY`).

---

## Acceptance criteria

* CI pipeline green; coverage gate passes.
* No cache restore errors on macOS.
* Commit signed **Jamal Al‑Sarraf <jalsarraf0@gmail.com>**.

---

MIT © 2025 Jamal Al‑Sarraf
