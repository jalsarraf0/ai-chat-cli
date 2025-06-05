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
