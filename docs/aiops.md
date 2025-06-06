# AI-Ops Helper

`ai-chat aiops watch --config <file>` streams logs from stdin and prints anomaly alerts based on regex patterns defined in the YAML config:

```yaml
patterns:
  - ERROR
  - FATAL
```

Lines matching any pattern are echoed with the pattern name. Use this to monitor logs and surface issues quickly.
