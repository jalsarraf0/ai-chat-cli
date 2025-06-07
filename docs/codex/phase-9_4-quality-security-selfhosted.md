<!--
AI-Chat-CLI • Codex Prompt
Phase 9.4 (v2) – Quality & Security on self-hosted “quality” runner
Save as: docs/codex/phase-9_4-quality-security-selfhosted.md
Generated: 2025-06-05
-->

# Phase 9.4 (v2) — **Quality & Security sweep on self-hosted runner** 🔒🛠
_The quality gate must run **only** on the Docker-based runner that carries
`quality`._

---

## 0 • Context / delta

* We now have **`amarillo-runner-01`** with labels
  `self-hosted`, `Linux`, `X64`, `quality`.
* Previous 9.4 prompt targeted `ubuntu-latest`; update required.

---

## 1 • CI workflow patch (`.github/workflows/ci.yml`)

```diff
@@
-  quality:
-    needs: test
-    runs-on: ubuntu-latest
+  quality:
+    needs: test
+    # target our dedicated container runner
+    runs-on: [self-hosted, linux, quality]
@@
-      - uses: actions/checkout@v4
+      - uses: actions/checkout@v4        # repo already present on runner
@@
```

> *`runs-on` array **must include** the custom label `quality` to avoid
>   scheduling on generic self-hosted runners.*

---

## 2 • Makefile / scripts – **no changes**
The tools bootstrap already succeeds inside the container.

---

## 3 • Acceptance checklist (self-hosted edition)

| Item | Pass criteria |
|------|---------------|
| **Runner picked** | `quality` job shows `amarillo-runner-01` in the Actions log header. |
| **Lint** | `golangci-lint run ./...` exits 0. |
| **Staticcheck** | `staticcheck ./...` exits 0. |
| **Security** | `gosec` + `govulncheck` report **0 medium+** findings. |
| **Coverage** | `make coverage-gate` ≥ 92 %. |
| **Matrix** | `test → quality → docs → snapshot` all green. |

---

## 4 • Tips for troubleshooting

* If the job queues indefinitely, verify the runner is **Online** and still
  carries the `quality` tag.
* To re-label an existing container runner:
  ```bash
  docker exec -it <container> bash
  config.sh remove --token <removeToken>
  config.sh --url https://github.com/jalsarraf0/ai-chat-cli             --token <regToken> --labels "self-hosted,Linux,quality" --unattended
  ```

---

MIT © 2025 Jamal Al-Sarraf
