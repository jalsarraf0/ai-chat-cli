# Plugin Development Tutorial

This guide walks you through creating an external plugin for AI-Chat-CLI.
Plugins extend the CLI without modifying its core.
For the full specification, see
[Phase 9 instructions](../codex/phase-9-plugin-system.md).

## 1. Create a Go module

```bash
mkdir myplugin && cd myplugin
cat > go.mod <<'MOD'
module example.com/myplugin

go 1.24
MOD
```

Implement the `Plugin` interface described in the Phase 9 document, then build
with `buildmode=plugin`:

```bash
go build -buildmode=plugin -o myplugin.so
```

## 2. Sign the plugin

Use `cosign` to sign the binary and place `myplugin.so` and
`myplugin.so.sig` in your plugins directory.

## 3. Install and test

Move the signed plugin into the directory shown by `ai-chat --plugins-dir`
and run:

```bash
ai-chat hello
```

You should see your plugin command appear in `ai-chat --help` output.
