# {{ .Site.Title | title }}

> {{ .Site.Params.tagline | default "📣 Chat from your terminal with AI super-powers" }}

{{ .Site.Home.Content }}

## Features
{{ range (where .Site.RegularPages "Section" "features") }}
- **{{ .Title }}** — {{ .Summary }}
{{ end }}

## Quick Start

```bash
curl -sSfL https://github.com/jalsarraf0/ai-chat-cli/releases/latest/download/install.sh | bash
```

_Generated on {{ now.Format "2006-01-02" }} with ❤️ by Hugo._
