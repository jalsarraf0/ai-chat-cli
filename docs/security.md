# Security

Release artefacts are signed using [Cosign](https://github.com/sigstore/cosign) with GitHub OIDC. The CI workflow signs each archive and attaches an SBOM:

```bash
cosign sign --yes dist/*.tar.gz
cosign attach sbom --yes dist/*.tar.gz
```

Verify signatures locally with:

```bash
cosign verify --certificate-identity-regexp 'github.com/.+' dist/*.tar.gz
```

## Static Analysis

The CI pipeline runs two scanners:

1. **gosec** – detects common security issues.
2. **govulncheck** – checks dependencies against the Go vulnerability database.

Suppress individual warnings by adding `// #nosec` for gosec or `//govulncheck:ignore` for govulncheck above the offending line.
