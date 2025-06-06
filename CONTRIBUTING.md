# Contributing

Thank you for helping improve **AI‑Chat‑CLI**. All changes go through pull requests.

## CI Runner Policy

Our GitHub Actions pipeline uses a gate‑first approach:

1. **unit-linux** and **unit-windows** run on dedicated self-hosted runners tagged
   `[self-hosted, Linux, X64, quality]` and `[self-hosted, Windows, X64, quality]`.
2. **unit-macos** runs on `macos-latest` **after** the Linux and Windows jobs
   succeed.
3. Downstream jobs depend on `unit-linux` at minimum to enforce the gate-first
   order.

Always ensure your PR targets the `phase-10/codex-apply` branch and that CI is
green before requesting review.

