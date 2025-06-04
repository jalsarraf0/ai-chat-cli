<!--
AI‑Chat‑CLI • Codex Prompt
Phase 7 – LLM Integration (OpenAI Adapter)
Save this file as docs/codex/phase-7-llm.md
Author: Jamal Al‑Sarraf <jalsarraf0@gmail.com>
-->

# Phase 7 Prompt – LLM Integration (OpenAI Adapter) 🤖
*Runner matrix: **Linux [self‑hosted, linux] · Windows [self‑hosted, windows] · macOS (macos‑latest)** — coverage gate **≥ 88 %** (expected ~90 %).*

---

## Deliverables (must all be implemented)

| # | Component | Precise requirement |
|---|-----------|---------------------|
| 1 | **`pkg/llm` interface** | `Client.Completion(ctx, req)` → streaming `Stream`. Structures defined exactly as spec. |
| 2 | **OpenAI back‑end** (`llm/openai`) | Uses `/v1/chat/completions` with `stream=true`. Reads key from **`AI_CHAT_API_KEY`** or `config.yaml:api_key`. Base URL override `AICHAT_BASE_URL`. |
| 3 | **Mock back‑end** (`llm/mock`) | Deterministic Stream for tests (`mock.New("hi", "there")`). |
| 4 | **TUI streaming** | Goroutine pumps tokens into Bubble Tea via `llmTokenMsg`; spinner hides when complete. |
| 5 | **CLI wiring** | `ai-chat ask <prompt>` streams to stdout. Flags & config: `model`, `temperature`, `max_tokens`. |
| 6 | **Tests** | 100 % in `llm/mock`; integration test in openai pkg skipped when `AI_CHAT_API_KEY` unset. TUI stream test validates spinner off. Project coverage ≥ 90 %. |
| 7 | **Docs** | `docs/llm.md` + README flag table. |
| 8 | **Security** | Redact API key in logs, 30 s timeout (`AICHAT_TIMEOUT` override). |

---

## Makefile additions

```make
live-openai-test:
go test ./pkg/llm/openai -run Live -v
```

---

## CI

No new jobs; unit tests use mock, so pipeline offline‑safe.

---

## Acceptance checklist

* CI green on Linux/Windows self‑hosted, macOS‑latest.  
* `ai-chat ask` works with live key.  
* Coverage ≥ 88 %.  
* Docs merged; commit signed **Jamal Al‑Sarraf**.

---

MIT © 2025 Jamal Al‑Sarraf
