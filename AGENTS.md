# AI-Chat-CLI Agents

This repository includes automated agents used in continuous integration.

- **CI Bot** – runs tests, linters and security scanners on pull requests.
- **Release Bot** – builds multi-platform archives and signs them with Cosign.
- **Docs Bot** – deploys the documentation site from the `docs/` folder.

External APIs used:

- **GitHub Actions** for CI/CD runners.
- **Sigstore** via Cosign and Rekor for artifact signing and verification.
- **OpenAI API** when running live integration tests (not in CI).

All bots operate under the policies described in `docs/codex/addons/always-online.md`.
