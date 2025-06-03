# Repository Guide for Codex

This file provides rules for Codex agents when modifying this repository.
It applies to the entire tree unless overridden by a nested `AGENTS.md`.

## 1. Foundation
- Go version: **1.24.3** (`CGO_ENABLED=0`).
- Module path: `github.com/jalsarraf0/ai-chat-cli`.
- Primary libraries:
  - `spf13/cobra` for CLI entry.
  - `charmbracelet/bubbletea` for TUI.
  - `spf13/viper` for configuration.
  - `charmbracelet/lipgloss` for styling.
  - `embed` for static assets.
  - `net/http` and `encoding/json` for API calls.
- Target shells: bash, zsh, fish, PowerShell (>=5), macOS zsh.
- Target OS: linux amd64/arm64, darwin amd64/arm64, windows amd64.

## 2. CI/CD & Tooling
- **Do not** modify GitHub Action workflows.
- The pipeline checks coverage (>=80%), `gosec`, and docs generation.
- Local helpers in `Makefile` or `scripts/` must run offline and honour
  `GOFLAGS=-mod=vendor` when set.
- Unit tests must compile with `-race` and keep total coverage >=81%.

## 3. Coding Standards
- Use `context.Context` in all exported functions performing I/O.
- Return `error` as the final value and wrap errors with `%w`.
- Use functional options for configuration (e.g. `WithAPIKey`).
- Table-driven tests with `t.Parallel()`.
- No globals except `var Version string` populated via `-ldflags`.
- CLI commands live under `cmd/ai-chat/`; internal packages under `internal/`.
- Imports ordered: standard library, third-party, local. Run `gofumpt` and
  `goimports` on Go files.

## 4. Branch, Commit & PR Hygiene
- Branch name format: `codex/{feature}`.
- First commit line <=72 chars, imperative present tense.
- Every PR must:
  1. Update or create unit tests.
  2. Update `docs/` if the public interface changes.
  3. Pass all CI jobs or remain a draft.
- Add the `codex-generated` label to each automated PR.

## 5. Shell Guidance
- Installation banners:
  - bash/zsh: `echo 'eval "$(ai-chat completion bash)"' >> ~/.bashrc`
  - fish: `ai-chat completion fish | tee ~/.config/fish/completions/ai-chat.fish`
  - PowerShell: ``ai-chat completion ps1 | Out-File -Encoding ascii $PROFILE``
- Detect the current shell via `$SHELL`, `ComSpec`, or `$env:PSExecutionPolicy`.
- Use ANSI colour via `github.com/muesli/termenv`; disable when `NO_COLOR`.

## 6. Release Engineering
- `make release` runs `goreleaser` locally (not in CI).
- Embed version info with `-ldflags "-s -w -X 'main.Version=$(git describe --tags --dirty)'"`.
- Generate completions and man pages in `dist/` during release builds.

## 7. Security & Compliance
- Do not hard-code tokens; read from the `AI_CHAT_API_KEY` environment variable
  or configuration file.
- Prefer `net/http` with `context` over third-party REST clients.
- Avoid reflection and the `unsafe` package.
- Keep third-party dependencies minimal and prefer the standard library.

## 8. Documentation
- Public packages require full Godoc comments; `gomarkdoc` verifies this.
- Update `docs/api.md` via the `gomarkdoc` target when interfaces change.
- Provide runnable code examples for every exported type using `// Output:`
  blocks.

## 9. Prompt & Chat Examples
- Store prompt templates under `embed/prompts/*.tmpl`.
- Examples must be shell-agnostic (use `$ ai-chat …`).
- At least one integration test must mock an LLM response.

## 10. Forbidden Practices
- Never alter `.github/workflows`.
- Avoid `sudo`, `curl | bash`, or non-Go build scripts in CI examples.
- Do not write files outside this repository or `$HOME/.config/ai-chat`.
- Do not commit generated binaries (except those in `dist/` produced by `goreleaser`).

## How to Use These Instructions
Always cross-check changes against these rules before finalising a PR. If any
instruction conflicts with a direct request from a maintainer, seek
clarification rather than guessing.

