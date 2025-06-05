# Releases

This project uses [Goreleaser](https://goreleaser.com/) for cross-platform builds and automated publishing.
Snapshot builds attach to pull requests while tagged commits trigger a full release.

Docker images are published to [ghcr.io/jalsarraf0/ai-chat-cli](https://ghcr.io/jalsarraf0/ai-chat-cli).
All archives and the Docker image are built for **amd64** only. Support for arm64 may return once the GolangCI-Lint issue is resolved.
Homebrew users can install via the `jalsarraf0/homebrew-ai-chat-cli` tap.
