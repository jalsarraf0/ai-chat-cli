# AGENTS Instructions

This repository contains Go code and documentation built with mdbook.

## Security
- Run `npm install` before doc generation and ensure `npm audit` reports **0 vulnerabilities**.
- If vulnerabilities appear, downgrade dependencies (e.g., `mdbook@1.4.1`) or run `npm audit fix --force` until clean.
- Do **not** commit the `node_modules` directory.
- Verify licensing via `addlicense -check $(git ls-files '*.go')`.

## Docs
- Build docs with `make docs` which runs `npm ci && npx mdbook build docs`.
